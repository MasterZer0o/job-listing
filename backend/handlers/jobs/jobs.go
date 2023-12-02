package handlers

import (
	// "encoding/json"
	"fmt"
	"log/slog"
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
	ExpLevel        string        `json:"level"`
	Title           string        `json:"title"`
	SalaryFrom      string        `json:"salaryFrom"`
	SalaryTo        string        `json:"salaryTo"`
	Currency        string        `json:"currency"`
	PostedAt        time.Time     `json:"postedAt"`
	SkillsJSON      string        `json:"-"`
	Skills          []string      `json:"skills"`
	EmploymentTypes []string      `json:"empTypes"`
	WorkModes       []string      `json:"workModes"`
	WorkTypes       []string      `json:"workTypes"`
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
	slog.Info("filters", slog.Any("", params))

	page := params["p"]
	cid := params["cid"]

	filters := ""
	for k := range params {
		switch k {
		case "tow":
		case "exp":
		case "emp":
		case "mode":
		case "tech":

		}
	}

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

	query := fmt.Sprintf(`SELECT jobs.id, location, title, jobs.created_at, remote_available,salary_from, salary_to,
	companies.name, companies.image, currencies.short_name, experience_levels.name,
json_agg(DISTINCT tech_skills.name),
json_agg(DISTINCT work_modes.name),
json_agg(DISTINCT employment_types.name),
json_agg(DISTINCT types_of_work.name)
FROM jobs
JOIN job_tech_skills on job_tech_skills.job_id = jobs.id
JOIN tech_skills ON job_tech_skills.tech_skill_id = tech_skills.id

JOIN job_employment_types on job_employment_types.job_id = jobs.id
JOIN employment_types on job_employment_types.emp_type_id= employment_types.id

JOIN job_experience_levels on job_experience_levels.job_id = jobs.id
JOIN experience_levels ON  job_experience_levels.exp_level_id = experience_levels.id

JOIN job_types_of_work on job_types_of_work.job_id = jobs.id
JOIN types_of_work ON  job_types_of_work.type_of_work_id = types_of_work.id

JOIN job_work_modes on job_work_modes.job_id = jobs.id
JOIN work_modes ON  job_work_modes.work_modes_id = work_modes.id

LEFT JOIN companies on companies.id = jobs.company_id
LEFT JOIN currencies on jobs.currency_id = currencies.id
%s
GROUP by jobs.id ,companies.name,companies.image,currencies.short_name, experience_levels.name LIMIT 20`, filters)
	// TODO: create new seeders for new tables, remove unused handle WHERE from query params,
	var rows pgx.Rows

	if cid != "" {
		s := time.Now()

		rows, _ = db.DB.Query(ctx.Context(), query, cid)

		slog.Info("With cid, time: ", slog.Any("", time.Since(s)))
	} else {
		s := time.Now()

		rows, _ = db.DB.Query(ctx.Context(), query)
		slog.Info("With page, time", slog.Any("", time.Since(s)), "page", page)
	}

	var err error
	for rows.Next() {
		var result JobsResult
		err = rows.Scan(&result.Id, &result.Location, &result.Title, &result.PostedAt, &result.RemoteAvailable, &result.SalaryFrom, &result.SalaryTo, &result.Company.Name, &result.Company.Image, &result.Currency, &result.ExpLevel, &result.Skills, &result.WorkModes, &result.EmploymentTypes, &result.WorkTypes)

		if err != nil {
			slog.Error("Scan error: ", "err", err)
		}

		// if result.SkillsJSON != "[null]" {
		// 	json.Unmarshal([]byte(result.SkillsJSON), &result.Skills)
		// } else {
		// 	result.Skills = []string{}
		// }
		results = append(results, result)
	}
	rows.Close()
	return ctx.JSON(map[string]interface{}{
		"data": results,
		// "cid":  results[len(results)-1].Id,
	})
}

func GetJob(ctx *fiber.Ctx) error {
	return nil
}
