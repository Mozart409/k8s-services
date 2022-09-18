package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(Timer())

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}

func Timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
		// return stack error if exist
		return err
	}
}
