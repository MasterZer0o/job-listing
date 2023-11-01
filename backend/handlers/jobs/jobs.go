package handlers

import (
	"encoding/json"
	"fmt"
	"main/db"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CompanyResult struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type JobsResult struct {
	Id              string        `json:"id"`
	Location        string        `json:"location"`
	RemoteAvailable bool          `json:"remoteAvailable"`
	Company         CompanyResult `json:"company"`
	Level           string        `json:"level"`
	Title           string        `json:"title"`
	SalaryFrom      string        `json:"salaryFrom"`
	SalaryTo        string        `json:"salaryTo"`
	Currency        string        `json:"currency"`
	PostedAt        time.Time     `json:"postedAt"`
	SkillsJSON      string        `json:"-"`
	Skills          []string      `json:"skills"`
}

func GetCount(ctx *fiber.Ctx) error {
	r, _ := db.DB.Query(ctx.Context(), "SELECT COUNT(*) FROM jobs")
	var count uint64
	r.Next()
	r.Scan(&count)
	r.Close()
	time.Sleep(time.Second)

	const PER_PAGE = 20
	var totalPages uint64
	if count%PER_PAGE == 0 {
		totalPages = count / PER_PAGE
	} else {
		totalPages = count/PER_PAGE + 1
	}

	return ctx.JSON(map[string]interface{}{
		"count":      count,
		"totalPages": totalPages,
	})
}

func GetJobs(ctx *fiber.Ctx) error {
	results := []JobsResult{}
	rows, _ := db.DB.Query(ctx.Context(),
		`SELECT jobs.id, location, level, title, jobs.created_at,remote_available, salary_from, salary_to,
	companies.name, companies.image, currencies.short_name, json_agg(job_skills.name) FROM jobs
	LEFT JOIN companies on companies.id = jobs.company_id
	LEFT JOIN currencies on jobs.currency_id = currencies.id
	LEFT JOIN job_skills on job_skills.job_id = jobs.id
	GROUP BY jobs.id, companies.name, companies.image, currencies.short_name LIMIT 20`)

	defer rows.Close()
	var err error
	for rows.Next() {
		var result JobsResult
		err = rows.Scan(&result.Id, &result.Location, &result.Level, &result.Title, &result.PostedAt, &result.RemoteAvailable, &result.SalaryFrom, &result.SalaryTo, &result.Company.Name, &result.Company.Image, &result.Currency, &result.SkillsJSON)

		if err != nil {
			fmt.Println(err)
		}

		if result.SkillsJSON != "[null]" {
			json.Unmarshal([]byte(result.SkillsJSON), &result.Skills)
		} else {
			result.Skills = []string{}
		}
		results = append(results, result)
	}

	return ctx.JSON(results)
}

func GetJob(ctx *fiber.Ctx) error {
	return nil
}
