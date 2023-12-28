package main

import (
	"log/slog"
	db "main/db"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"

	"github.com/joho/godotenv"
	"gitlab.com/greyxor/slogor"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Network:               "tcp",
	})

	app.Use(helmet.New())

	slog.SetDefault(slog.New(slogor.NewHandler(os.Stdout, &slogor.Options{
		TimeFormat: time.TimeOnly,
		Level:      slog.LevelDebug,
		ShowSource: true,
	})))

	godotenv.Load()

	db.Connect()

	allowedOrigin, ok := os.LookupEnv("ALLOWED_ORIGIN")

	if !ok {
		slog.Error("ALLOWED_ORIGIN ENV not set")
		os.Exit(1)
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigin,
		AllowCredentials: true,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	SetupRoutes(app)

	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		slog.Error("PORT ENV not set")
		os.Exit(1)
	}

	slog.Info("Server started on port " + PORT)
	app.Listen(":" + PORT)
}
