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
	"github.com/gofiber/swagger"

	"go.uber.org/zap"

	_ "github.com/cjtim/be-friends-api/docs"
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

// @title Be Friends API
// @version 1.0

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @BasePath /
func prepareFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandling,
		BodyLimit:    100 * 1024 * 1024, // Limit file size to 100MB
	})

	app.Get("/api/swagger/*", swagger.HandlerDefault)
	app.Get("/api/swagger/*", swagger.New(swagger.Config{
		URL:          "/api/doc.json",
		DeepLinking:  true,
		DocExpansion: "none",
	}))

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
