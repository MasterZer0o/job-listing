package main

import (
	jobs "main/handlers/jobs"
	user "main/handlers/user"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	userGroup := app.Group("/user")

	userGroup.Post("/register", user.Register)
	userGroup.Post("/login", user.Login)
	userGroup.Get("/auth", user.Authenticate)
	userGroup.Post("/logout", user.Logout)

	app.Get("/jobs", jobs.GetJobs)
	app.Get("/job/:id", jobs.GetJob)

}