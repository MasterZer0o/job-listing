package handlers

import (
	"errors"
	"log/slog"
	"main/db"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Email      string
	Password   string
	Remember   bool
	CheckSaved bool
}

type foundUser struct {
	Id       uint
	Email    string
	Password string
	IsSaved  bool
}

func Login(ctx *fiber.Ctx) error {
	loginData := loginBody{}
	ctx.BodyParser(&loginData)

	var query string
	var err error
	user := foundUser{}
	if !loginData.CheckSaved {
		query = "SELECT email, password, id FROM users WHERE email = $1"
		err = db.DB.QueryRow(ctx.Context(), query, loginData.Email).Scan(&user.Email, &user.Password, &user.Id)
	} else {
		query = `SELECT email, password, users.id, saved_jobs.job_id IS NOT NULL is_saved
		FROM users
		LEFT JOIN saved_jobs ON saved_jobs.user_id = users.id
		WHERE email = $1`
		err = db.DB.QueryRow(ctx.Context(), query, loginData.Email).Scan(&user.Email, &user.Password, &user.Id, &user.IsSaved)
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return ctx.JSON(fiber.Map{
			"error": "Email or password is incorrect",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": "Email or password is incorrect",
		})
	}

	sessionId := uuid.New()
	_, err = db.DB.Exec(ctx.Context(), "INSERT into sessions (id, user_id) VALUES ($1, $2)", sessionId, user.Id)
	// cookie := &fiber.Cookie{
	// 	Name:     "session",
	// 	HTTPOnly: true,
	// 	Value:    sessionId.String(),
	// 	Secure:   true,
	// 	SameSite: "none",
	// 	// Domain:   "vercel.app",
	// }

	if err != nil {
		slog.Error("Failed to create user session in db.", "err", err)
	}

	// if loginData.Remember {
	// 	cookie.MaxAge = 2590000
	// } else {
	// 	cookie.SessionOnly = true
	// }

	// ctx.Cookie(cookie)
	if !loginData.CheckSaved {
		return ctx.JSON(map[string]interface{}{
			"sid": sessionId.String(),
		})
	}
	return ctx.JSON(map[string]interface{}{
		"sid": sessionId.String(),
		"saved": user.IsSaved,
	})

}
