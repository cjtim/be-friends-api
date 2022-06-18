package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/handlers"
	"github.com/cjtim/be-friends-api/handlers/middlewares"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	err := configs.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Initial Global logger
	logger := middlewares.InitZap()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Connect DB
	closeFn, err := repository.PrepareDB()
	defer closeFn() // Close all DB conn
	if err != nil {
		zap.L().Error("prepare db error", zap.Error(err))
		return 1
	}

	// Prepare API route
	app := prepareFiber()
	setupCloseHandler(app)

	// Start accept connection
	listen := fmt.Sprintf(":%d", configs.Config.Port)
	err = app.Listen(listen)
	if err != nil {
		zap.L().Error("fiber error", zap.Error(err))
		return 1
	}
	return 0
}

func prepareFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandling,
		BodyLimit:    100 * 1024 * 1024, // Limit file size to 100MB
	})
	app.Use(middlewares.Cors())
	app.Use(middlewares.RequestLog())
	handlers.Route(app) // setup router path
	return app
}

// setupCloseHandler - What to do when got ctrl+c SIGTERM
func setupCloseHandler(app *fiber.App) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		zap.L().Info("Got SIGTERM, terminating program...")
		app.Server().Shutdown()
	}()
}
