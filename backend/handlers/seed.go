package handlers

import (
	"fmt"
	"strconv"

	"main/db/seeders"

	"github.com/gofiber/fiber/v2"
)

func SeedHandler(ctx *fiber.Ctx) error {
	parameters := ctx.Queries()
	seedTarget := parameters["target"]
	count := parameters["count"]

	if seedTarget == "" || count == "" {
		return ctx.JSON(map[string]interface{}{
			"success": false,
			"error":   "Count or target query param not provided",
		})
	}
	countInt, _ := strconv.ParseInt(count, 10, 64)

	switch seedTarget {
	case "jobs":
		seeders.SeedJobs(int(countInt), parameters["reset"] == "true")

	case "skills":
		seeders.SeedSkills()
	case "users":

	default:
		return ctx.JSON(map[string]interface{}{
			"success": false,
			"message": fmt.Sprintf("Target %s unknown", seedTarget),
		})

	}
	return ctx.JSON(map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Successfully inserted %d rows into %s", countInt, seedTarget),
	})
}
