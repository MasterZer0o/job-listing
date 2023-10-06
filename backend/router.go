package main

import (
	"main/handlers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	user := app.Group("/user")

	user.Post("/register", handlers.Register)

}
