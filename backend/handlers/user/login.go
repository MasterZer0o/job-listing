package handlers

import (
	"errors"
	"fmt"
	"main/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Email    string
	Password string
	Remember bool
}

type foundUser struct {
	Id       uint
	Email    string
	Password string
}

func Login(ctx *fiber.Ctx) error {
	loginData := loginBody{}
	err := ctx.BodyParser(&loginData)

	if err != nil {
		fmt.Println("Body parse error: ", err)
	}

	time.Sleep(1 * time.Second)

	user := foundUser{}
	err = db.DB.QueryRow(ctx.Context(), "SELECT email, password, id FROM users WHERE email = $1",
		loginData.Email).Scan(&user.Email, &user.Password, &user.Id)

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
	cookie := &fiber.Cookie{
		Name:     "session",
		HTTPOnly: true,
		Value:    sessionId.String(),
		Secure:   true,
		SameSite: "none",
	}
	// TODO: return sessionId in response instead of cookie
	if loginData.Remember {
		cookie.MaxAge = 2590000
	} else {
		cookie.SessionOnly = true
	}

	ctx.Cookie(cookie)

	return ctx.SendStatus(200)
}