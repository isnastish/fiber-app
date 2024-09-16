package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/isnastish/fiber-app/pkg/log"
)

func main() {
	config := fiber.Config{
		Prefork: true,
		// DisableStartupMessage: true, // (should be disable in production)
		// TODO: Clarify what strict routing means
		StrictRouting: true,
	}

	app := fiber.New(config)

	app.Get("/api/v1/books", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello from fiber")
	})

	osSigChan := make(chan os.Signal, 1)
	signal.Notify(osSigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen("0.0.0.0:5000"); err != nil {
			log.Logger.Fatal("Failed to listen %v", err)
		}
	}()

	<-osSigChan
	// gracefully shutdown the server
	app.Shutdown()
}
