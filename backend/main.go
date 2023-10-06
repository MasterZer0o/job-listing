package main

import (
	"fmt"
	db "main/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	godotenv.Load()
	db.Connect()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	Router(app)

	fmt.Print("Server started on port 5000.\n")
	app.Listen("localhost:5000")
}
