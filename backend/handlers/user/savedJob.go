package handlers

import (
	"log/slog"
	"main/db"

	"github.com/gofiber/fiber/v2"
)

type body struct {
	JobId string `json:"jobId"`
}

func SaveJob(ctx *fiber.Ctx) error {
	var data body
	ctx.BodyParser(&data)

	slog.Info(ctx.Cookies("session"))
	userId := getUserId(ctx, ctx.Cookies("session"))

	db.DB.Exec(ctx.Context(), "INSERT INTO saved_jobs (user_id, job_id) VALUES ($1,$2)", userId, data.JobId)

	return ctx.SendStatus(200)
}

func getUserId(ctx *fiber.Ctx, sessionId string) string {
	row := db.DB.QueryRow(ctx.Context(), "SELECT user_id from sessions WHERE id = $1", sessionId)
	var userId string
	err := row.Scan(&userId)

	if err != nil {
		slog.Error("Scan error", "err", err)
	}

	return userId
}
