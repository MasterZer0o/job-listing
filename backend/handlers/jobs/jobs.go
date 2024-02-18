package handlers

import (
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
type countResponse struct {
	Count      uint64 `json:"count"`
	TotalPages uint64 `json:"totalPages"`
}

func GetCount(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	var filters []string

	sqlQuery := "SELECT COUNT(distinct jobs.id) FROM jobs"

	for k := range params {
		switch k {
		case "tech":
			sqlQuery += ` LEFT JOIN job_tech_skills ON job_tech_skills.job_id = jobs.id
			LEFT JOIN tech_skills ON job_tech_skills.tech_skill_id = tech_skills.id`
			filters = append(filters, "job_tech_skills.tech_skill_id IN ("+params["tech"]+")")

		case "tow":
			sqlQuery += ` LEFT JOIN job_types_of_work ON job_types_of_work.job_id = jobs.id
			LEFT JOIN types_of_work ON job_types_of_work.type_of_work_id = types_of_work.id`
			filters = append(filters, "job_types_of_work.type_of_work_id IN ("+params["tow"]+")")

		case "exp":
			sqlQuery += ` LEFT JOIN job_experience_levels ON job_experience_levels.job_id = jobs.id
			LEFT JOIN experience_levels ON job_experience_levels.exp_level_id = experience_levels.id`
			filters = append(filters, "job_experience_levels.exp_level_id IN ("+params["exp"]+")")
		case "emp":
			sqlQuery += ` LEFT JOIN job_employment_types ON job_employment_types.job_id = jobs.id
			LEFT JOIN employment_types ON job_employment_types.emp_type_id = employment_types.id`
			filters = append(filters, "job_employment_types.emp_type_id IN ("+params["emp"]+")")
		case "mode":
			sqlQuery += ` LEFT JOIN job_work_modes ON job_work_modes.job_id = jobs.id
			LEFT JOIN work_modes ON job_work_modes.work_mode_id = work_modes.id`
			filters = append(filters, "job_work_modes.work_mode_id IN ("+params["mode"]+")")

		case "ms":
			filters = append(filters, "salary_from >="+params["ms"])

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
	r, _ := db.DB.Query(ctx.Context(), sqlQuery)
	var count uint64
	r.Next()
	r.Scan(&count)
	r.Close()

	const RESULTS_PER_PAGE = 20
	var totalPages uint64
	if count%RESULTS_PER_PAGE == 0 {
		totalPages = count / RESULTS_PER_PAGE
	} else {
		totalPages = count/RESULTS_PER_PAGE + 1
	}

	return ctx.JSON(countResponse{
		Count:      count,
		TotalPages: totalPages,
	})
}

func GetJobs(ctx *fiber.Ctx) error {
	params := ctx.Queries()

	page := params["p"]
	cid := params["cid"]

	var filters []string

	having := ""
	joins := ""
	for k := range params {
		switch k {
		case "tech":
			having = "HAVING array_agg(job_tech_skills.tech_skill_id) && ARRAY[" + params["tech"] + "]"

		case "tow":
			joins += ` LEFT JOIN job_types_of_work ON job_types_of_work.job_id = jobs.id
			LEFT JOIN types_of_work ON job_types_of_work.type_of_work_id = types_of_work.id`
			filters = append(filters, "job_types_of_work.type_of_work_id IN ("+params["tow"]+")")

		case "exp":
			joins += ` LEFT JOIN job_experience_levels ON job_experience_levels.job_id = jobs.id
			LEFT JOIN experience_levels ON job_experience_levels.exp_level_id = experience_levels.id`
			filters = append(filters, "job_experience_levels.exp_level_id IN ("+params["exp"]+")")
		case "emp":
			joins += ` LEFT JOIN job_employment_types ON job_employment_types.job_id = jobs.id
			LEFT JOIN employment_types ON job_employment_types.emp_type_id = employment_types.id`
			filters = append(filters, "job_employment_types.emp_type_id IN ("+params["emp"]+")")
		case "mode":
			joins += ` LEFT JOIN job_work_modes ON job_work_modes.job_id = jobs.id
			LEFT JOIN work_modes ON job_work_modes.work_mode_id = work_modes.id`
			filters = append(filters, "job_work_modes.work_mode_id IN ("+params["mode"]+")")
		case "ms":
			filters = append(filters, "salary_from >="+params["ms"])
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
	array_agg(distinct tech_skills.name) AS tech_skills, location,level, title, jobs.created_at, remote_available,salary_from, salary_to,
	companies.name, companies.image, currencies.short_name
FROM jobs
LEFT JOIN companies on companies.id = jobs.company_id
LEFT JOIN currencies on jobs.currency_id = currencies.id 
LEFT JOIN job_tech_skills ON job_tech_skills.job_id = jobs.id
LEFT JOIN tech_skills ON job_tech_skills.tech_skill_id = tech_skills.id
%s
%s
GROUP by jobs.id, companies.name, companies.image, currencies.short_name
%s
ORDER BY jobs.id 
LIMIT 20`, joins, filtersStr, having)

	var rows pgx.Rows
	var err error

	if cid != "" {
		s := time.Now()
		rows, err = db.DB.Query(ctx.Context(), query, cid)

		slog.Info("With cid", "time", time.Since(s))
	} else {
		s := time.Now()
		rows, err = db.DB.Query(ctx.Context(), query)
		slog.Info("With page", "time", time.Since(s), "page", page)
	}

	if err != nil {
		slog.Error("Query error", "err", err)
		return ctx.JSON(response{
			Error: "Failed to get data",
		})
	}
	results := []JobsResult{}

	for rows.Next() {
		var result JobsResult
		err = rows.Scan(&result.Id, &result.Skills, &result.Location, &result.ExpLevel, &result.Title, &result.PostedAt, &result.RemoteAvailable, &result.SalaryFrom, &result.SalaryTo, &result.Company.Name, &result.Company.Image, &result.Currency)

		if err != nil {
			slog.Error("Scan error: ", "err", err)
			return ctx.JSON(response{
				Error: "Failed to get data",
			})
		}
		results = append(results, result)
	}
	rows.Close()

	var newCid string
	if len(results) > 0 {
		newCid = results[len(results)-1].Id
	}

	return ctx.JSON(response{
		Results: results,
		Cid:     newCid,
	})
}

type response struct {
	Results []JobsResult `json:"data"`
	Cid     string       `json:"cid,omitempty"`
	Error   string       `json:"error,omitempty"`
}
