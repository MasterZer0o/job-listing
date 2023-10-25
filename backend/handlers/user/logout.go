package handlers

import (
	"main/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"
)

func Logout(ctx *fiber.Ctx) error {
	sessionId := ctx.Request().Header.Cookie("session")
	db := db.DB

	go func(db *pgxpool.Pool, sessionId string, ctx *fasthttp.RequestCtx) {
		db.Exec(ctx, "DELETE FROM sessions WHERE id = $1", sessionId)
	}(db, string(sessionId), ctx.Context())

	ctx.Cookie(&fiber.Cookie{
		Name:     "session",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "none",
		Expires:  time.Unix(0, 0),
	})
	return ctx.SendStatus(200)
}
