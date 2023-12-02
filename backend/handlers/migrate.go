package handlers

import (
	"fmt"
	"main/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Migrate(ctx *fiber.Ctx) error {
	v := ctx.Queries()
	fmt.Println(v["v"])
	var forceVer int
	if v["v"] != "" {
		forceVer, _ = strconv.Atoi(v["v"])
		db.Migrate(&forceVer)

	} else {
		db.Migrate(nil)
	}
	return ctx.JSON(map[string]interface{}{
		"success": "true",
	})
}
