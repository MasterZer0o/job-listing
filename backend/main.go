package main

import (
	"fmt"
	"log/slog"
	db "main/db"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gitlab.com/greyxor/slogor"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Network:               "tcp",
	})

	godotenv.Load()

	db.Connect()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	slog.SetDefault(slog.New(slogor.NewHandler(os.Stdout, &slogor.Options{
		TimeFormat: time.TimeOnly,
		Level:      slog.LevelDebug,
		ShowSource: true,
	})))

	Router(app)

	fmt.Print("Server started on port 5000.\n")
	app.Listen(":5000")
}
