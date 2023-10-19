package handlers

import (
	"errors"
	"main/db"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type userSession struct {
	SessionId string
	UserId    uint
	CreatedAt string
}

func Authenticate(ctx *fiber.Ctx) error {
	sessionId := ctx.Request().Header.Cookie("session")
	var err error
	user := userSession{}

	if len(sessionId) == 0 {
		return ctx.JSON(fiber.Map{
			"success": false,
		})
	}

	err = db.DB.QueryRow(ctx.Context(), "SELECT user_id, id, created_at FROM sessions WHERE id = $1", string(sessionId)).
		Scan(&user.UserId, &user.SessionId, &user.CreatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return ctx.JSON(fiber.Map{
			"success": false,
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"userId":  user.UserId,
	})

}
