package handlers

import (
	"context"
	"log/slog"
	"main/db"

	"github.com/gofiber/fiber/v2"
)

type body struct {
	JobId   string `json:"jobId"`
	IsSaved bool   `json:"isSaved"`
}

func SaveOrRemoveJob(ctx *fiber.Ctx) error {
	var data body
	ctx.BodyParser(&data)

	userId := getUserId(ctx, ctx.Cookies("session"))
	var err error

	if data.IsSaved {
		err = removeSavedJob(userId, data.JobId)
	} else {
		err = saveJob(userId, data.JobId)
	}

	if err != nil {
		slog.Error("Save or remove saved job error", "err", err)
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}

func saveJob(userId string, jobId string) error {
	_, err := db.DB.Exec(context.Background(), "INSERT INTO saved_jobs (user_id, job_id) VALUES ($1,$2)", userId, jobId)
	return err
}

func removeSavedJob(userId string, jobId string) error {
	_, err := db.DB.Exec(context.Background(), "DELETE FROM saved_jobs WHERE user_id=$1 AND job_id=$2", userId, jobId)
	return err
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
