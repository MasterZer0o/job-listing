package main

import (
	"main/handlers"
	jobs "main/handlers/jobs"
	user "main/handlers/user"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userGroup := app.Group("/user")

	userGroup.Post("/register", user.Register)
	userGroup.Post("/login", user.Login)
	userGroup.Get("/auth", user.Authenticate)
	userGroup.Post("/logout", user.Logout)
	userGroup.Post("/saved-jobs", user.SaveJob)

	app.Get("/jobs", jobs.GetJobs)
	app.Get("/jobs/count", jobs.GetCount)
	app.Get("/job/:id", jobs.GetJob)

	isDev, _ := os.LookupEnv("DEVELOPMENT")

	if isDev == "true" {
		app.Get("/db/seed", handlers.SeedHandler)
		app.Get("/db/migrate", handlers.Migrate)
	}

}
