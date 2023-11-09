package handlers

import (
	"encoding/json"
	"fmt"
	"main/db"

	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
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
	params := ctx.Queries()
	page := params["p"]
	cid := params["cid"]

	filters := ""
	if cid != "" {
		filters = "WHERE jobs.id > $1"
	}

	if cid == "" && page != "" {
		const JOBS_PER_PAGE = 20
		pageNumber, _ := strconv.Atoi(page)
		filters += fmt.Sprintf(`WHERE jobs.id >= (SELECT MIN(jobs.id) FROM jobs WHERE jobs.id > %d)
		AND jobs.id <= (SELECT MAX(jobs.id) FROM jobs WHERE jobs.id > %d)`,
			(pageNumber-1)*JOBS_PER_PAGE, (pageNumber-1)*JOBS_PER_PAGE)
	}

	query := fmt.Sprintf(`SELECT jobs.id, location, level, title, jobs.created_at,remote_available, salary_from, salary_to,
	companies.name, companies.image, currencies.short_name, json_agg(job_skills.name) FROM jobs
	LEFT JOIN companies on companies.id = jobs.company_id
	LEFT JOIN currencies on jobs.currency_id = currencies.id
	LEFT JOIN job_skills on job_skills.job_id = jobs.id
	%s
	GROUP BY jobs.id, companies.name, companies.image, currencies.short_name ORDER BY jobs.id LIMIT 20`, filters)

	var rows pgx.Rows

	if cid != "" {
		s := time.Now()

		rows, _ = db.DB.Query(ctx.Context(), query, cid)

		fmt.Println("With cid time: ", time.Since(s))
	} else {
		s := time.Now()
		rows, _ = db.DB.Query(ctx.Context(), query)
		fmt.Printf("Page %s time: %s\n", page, time.Since(s))
	}

	var err error
	for rows.Next() {
		var result JobsResult
		err = rows.Scan(&result.Id, &result.Location, &result.Level, &result.Title, &result.PostedAt, &result.RemoteAvailable, &result.SalaryFrom, &result.SalaryTo, &result.Company.Name, &result.Company.Image, &result.Currency, &result.SkillsJSON)

		if err != nil {
			fmt.Println("Scan error: ", err)
		}

		if result.SkillsJSON != "[null]" {
			json.Unmarshal([]byte(result.SkillsJSON), &result.Skills)
		} else {
			result.Skills = []string{}
		}
		results = append(results, result)
	}
	rows.Close()
	return ctx.JSON(map[string]interface{}{
		"data": results,
		"cid":  results[len(results)-1].Id,
	})
}

func GetJob(ctx *fiber.Ctx) error {
	return nil
}
