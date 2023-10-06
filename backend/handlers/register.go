package handlers

import

// import (
// 	"errors"
// 	"main/models"

// "github.com/go-playground/validator/v10"
"github.com/gofiber/fiber/v2"

// 	"golang.org/x/crypto/bcrypt"
// )

// type RequestData struct {
// 	Email          string `validate:"email"`
// 	Password       string `validate:"gte=6,eqfield=PasswordRepeat"`
// 	PasswordRepeat string `validate:"gte=6"`
// }

func Register(ctx *fiber.Ctx) error {
	// 	incomingData := new(RequestData)
	// 	var err error
	// 	if err = ctx.BodyParser(incomingData); err != nil {
	// 		return err
	// 	}

	// 	err = incomingData.validate()

	// 	if err != nil {
	// 		ctx.Status(406)
	// 		return ctx.JSON(fiber.Map{
	// 			"error": err.Error(),
	// 		})
	// 	}

	// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(incomingData.Password), 8)

	// 	if err != nil {
	// 		return ctx.SendStatus(500)
	// 	}

	// 	data := models.User{
	// 		Email:    incomingData.Email,
	// 		Password: string(hashedPassword),
	// 	}

	// 	result := config.Db.Create(&data)

	// 	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
	// 		return ctx.JSON(fiber.Map{
	// 			"error": "This email is already registered.",
	// 		})
	// 	}

	// 	return ctx.SendStatus(202)
	// }

	// var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

	// func (data *RequestData) validate() error {
	// 	if err := validate.Struct(data); err != nil {
	// 		var message string

	// 		switch err.(validator.ValidationErrors)[0].Tag() {
	// 		case "eqfield":
	// 			message = "Passwords are not equal"
	// 		case "gte":
	// 			message = "Password is too short"
	// 		case "email":
	// 			message = "Email is invalid"
	// 		}
	// 		return errors.New(message)
	// 	}

	return nil
}
