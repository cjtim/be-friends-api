package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cjtim/be-friends-api/configs"
	controller "github.com/cjtim/be-friends-api/internal/app/controllers"
	"github.com/cjtim/be-friends-api/internal/app/middlewares"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	logger := middlewares.InitZap()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	err := repository.CreateConnection()
	if err != nil {
		return 1
	}

	app := startServer()
	setupCloseHandler(app)

	listen := fmt.Sprintf(":%d", configs.Config.Port)
	if err := app.Listen(listen); err != nil {
		zap.L().Error("fiber start error", zap.Error(err))
		return 1
	}
	return 0
}

func startServer() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandling,
		BodyLimit:    100 * 1024 * 1024, // Limit file size to 100MB
	})
	app.Use(middlewares.Cors())
	app.Use(middlewares.RequestLog())
	controller.Route(app) // setup router path
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
		// repository.Client.Disconnect()
		app.Server().Shutdown()
	}()
}
