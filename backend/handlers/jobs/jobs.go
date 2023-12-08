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
	Skills          []string      `json:"skills"`
}
// TODO: add "minimum salary filter"

func GetCount(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	var filters []string

	sqlQuery := "SELECT COUNT(*) FROM jobs"

	for k := range params {
		switch k {
		case "tech":
			sqlQuery += ` LEFT JOIN (
				SELECT job_id, array_agg(tech_skills.id) AS skills
				FROM job_tech_skills
				LEFT JOIN tech_skills ON job_tech_skills.tech_skill_id = tech_skills.id
				GROUP BY job_id
				) AS skills ON skills.job_id = jobs.id`
			filters = append(filters, "skills && "+"ARRAY["+params["tech"]+"]::smallint[]")

		case "tow":
			sqlQuery += ` LEFT JOIN (
				SELECT job_id, array_agg(types_of_work.id) AS tow
				FROM job_types_of_work
				LEFT JOIN types_of_work ON job_types_of_work.type_of_work_id = types_of_work.id
				GROUP BY job_id
				) AS tow ON tow.job_id = jobs.id`
			filters = append(filters, "tow && "+"ARRAY["+params["tow"]+"]::smallint[]")

		case "exp":
			sqlQuery += ` LEFT JOIN (
				SELECT job_id, array_agg(experience_levels.id) AS exp
				FROM job_experience_levels
				LEFT JOIN experience_levels ON job_experience_levels.exp_level_id = experience_levels.id
				GROUP BY job_id
				) AS exp ON exp.job_id = jobs.id`
			filters = append(filters, "exp && "+"ARRAY["+params["exp"]+"]::smallint[]")
		case "emp":
			sqlQuery += ` LEFT JOIN (
				SELECT job_id, array_agg(employment_types.id) AS emp
				FROM job_employment_types
				LEFT JOIN employment_types ON job_employment_types.emp_type_id = employment_types.id
				GROUP BY job_id
				) AS emp ON emp.job_id = jobs.id`
			filters = append(filters, "emp && "+"ARRAY["+params["emp"]+"]::smallint[]")
		case "mode":
			sqlQuery += ` LEFT JOIN (
				SELECT job_id, array_agg(work_modes.id) AS mode
				FROM job_work_modes
				LEFT JOIN work_modes ON job_work_modes.work_mode_id = work_modes.id
				GROUP BY job_id
				) AS mode ON mode.job_id = jobs.id`
			filters = append(filters, "mode && "+"ARRAY["+params["mode"]+"]::smallint[]")

		}
	}

	if len(filters) > 0 {
		sqlQuery += " WHERE"

		for idx, filterStr := range filters {
			sqlQuery += " " + filterStr

			if idx+1 != len(filters) {
				sqlQuery += " AND"
			}
		}
	}
	// slog.Info(sqlQuery)
	r, _ := db.DB.Query(ctx.Context(), sqlQuery)
	var count uint64
	r.Next()
	r.Scan(&count)
	r.Close()

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

	var filters []string
	var filterSelectFields string
	joins := ""
	for k := range params {
		switch k {
		case "tech":
			joins += ` LEFT JOIN (
				SELECT job_id, array_agg(tech_skills.id) AS skills
				FROM job_tech_skills
				LEFT JOIN tech_skills ON job_tech_skills.tech_skill_id = tech_skills.id
				GROUP BY job_id
				) AS skills ON skills.job_id = jobs.id`
			filters = append(filters, "skills && "+"ARRAY["+params["tech"]+"]::smallint[]")

		case "tow":
			joins += ` LEFT JOIN (
				SELECT job_id, array_agg(types_of_work.id) AS tow
				FROM job_types_of_work
				LEFT JOIN types_of_work ON job_types_of_work.type_of_work_id = types_of_work.id
				GROUP BY job_id
				) AS tow ON tow.job_id = jobs.id`
			filters = append(filters, "tow && "+"ARRAY["+params["tow"]+"]::smallint[]")

		case "exp":
			joins += ` LEFT JOIN (
				SELECT job_id, array_agg(experience_levels.id) AS exp
				FROM job_experience_levels
				LEFT JOIN experience_levels ON job_experience_levels.exp_level_id = experience_levels.id
				GROUP BY job_id
				) AS exp ON exp.job_id = jobs.id`
			filters = append(filters, "exp && "+"ARRAY["+params["exp"]+"]::smallint[]")
		case "emp":
			joins += ` LEFT JOIN (
				SELECT job_id, array_agg(employment_types.id) AS emp
				FROM job_employment_types
				LEFT JOIN employment_types ON job_employment_types.emp_type_id = employment_types.id
				GROUP BY job_id
				) AS emp ON emp.job_id = jobs.id`
			filters = append(filters, "emp && "+"ARRAY["+params["emp"]+"]::smallint[]")
		case "mode":
			joins += ` LEFT JOIN (
				SELECT job_id, array_agg(work_modes.id) AS mode
				FROM job_work_modes
				LEFT JOIN work_modes ON job_work_modes.work_mode_id = work_modes.id
				GROUP BY job_id
				) AS mode ON mode.job_id = jobs.id`
			filters = append(filters, "mode && "+"ARRAY["+params["mode"]+"]::smallint[]")

		}
	}

	if cid != "" {
		filters = append(filters, "jobs.id > $1")
	}

	if cid == "" && page != "" {
		const JOBS_PER_PAGE = 20
		pageNumber, _ := strconv.Atoi(page)
		filters = append(filters, fmt.Sprintf(`jobs.id >= (SELECT MIN(jobs.id) FROM jobs WHERE jobs.id > %d)
		AND jobs.id <= (SELECT MAX(jobs.id) FROM jobs WHERE jobs.id > %d)`,
			(pageNumber-1)*JOBS_PER_PAGE, (pageNumber-1)*JOBS_PER_PAGE))
	}

	filtersStr := ""
	if len(filters) > 0 {
		filtersStr += "WHERE"

		for idx, filterStr := range filters {
			filtersStr += " " + filterStr

			if idx+1 != len(filters) {
				filtersStr += " AND"
			}
		}
	}

	query := fmt.Sprintf(`SELECT jobs.id, 
(SELECT json_agg(tech_skills.name) FROM job_tech_skills LEFT JOIN tech_skills ON job_tech_skills.tech_skill_id = tech_skills.id WHERE job_tech_skills.job_id = jobs.id), location, title, jobs.created_at, remote_available,salary_from, salary_to,
	companies.name, companies.image, currencies.short_name %s
FROM jobs
LEFT JOIN companies on companies.id = jobs.company_id
LEFT JOIN currencies on jobs.currency_id = currencies.id %s
%s
GROUP by jobs.id , companies.name, companies.image, currencies.short_name
ORDER BY jobs.id 
LIMIT 20`, filterSelectFields, joins, filtersStr)
	var rows pgx.Rows
	var err error
	slog.Info(query)
	if cid != "" {
		s := time.Now()
		rows, err = db.DB.Query(ctx.Context(), query, cid)

		if err != nil {
			slog.Error("Query error", "err", err)

		}

		slog.Info("With cid, time: ", slog.Any("", time.Since(s)))
	} else {
		s := time.Now()
		rows, _ = db.DB.Query(ctx.Context(), query)
		slog.Info("With page, time", slog.Any("", time.Since(s)), "page", page)
	}

	for rows.Next() {
		var result JobsResult
		err = rows.Scan(&result.Id, &result.Skills, &result.Location, &result.Title, &result.PostedAt, &result.RemoteAvailable, &result.SalaryFrom, &result.SalaryTo, &result.Company.Name, &result.Company.Image, &result.Currency)

		if err != nil {
			slog.Error("Scan error: ", "err", err)
		}
		results = append(results, result)
	}
	rows.Close()
	var newCid string
	if len(results) > 0 {
		newCid = results[len(results)-1].Id
	}

	return ctx.JSON(Response{
		Results: results,
		Cid:     newCid,
	})
}

type Response struct {
	Results []JobsResult `json:"data"`
	Cid     string       `json:"cid,omitempty"`
}

func GetJob(ctx *fiber.Ctx) error {
	return nil
}
