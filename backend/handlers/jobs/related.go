package handlers

import (
	"log/slog"
	"main/db"

	"github.com/gofiber/fiber/v2"
)

type RelatedJobSalary struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type RelatedJob struct {
	Id              string            `json:"id"`
	Title           string            `json:"title"`
	RemoteAvailable bool              `json:"remoteAvailable"`
	CompanyName     string            `json:"companyName"`
	IsSaved         bool              `json:"isSaved"`
	Skills          []string          `json:"skills"`
	Salary          *RelatedJobSalary `json:"salary"`
}

// TODO:
/*
-improve seeding -> do something with level, salary with level (?)
*/

func GetRelated(ctx *fiber.Ctx) error {
	jobId := ctx.Params("id")
	row := db.DB.QueryRow(ctx.Context(), "SELECT title FROM jobs WHERE id=$1", jobId)
	sessionId := ctx.Cookies("session")

	var s [3]string
	if sessionId != "" {
		row := db.DB.QueryRow(ctx.Context(), "SELECT user_id FROM sessions WHERE id=$1", sessionId)
		var userId string
		row.Scan(&userId)

		if userId != "" {
			s[0] = ", CASE WHEN sj.job_id=jobs.id THEN TRUE ELSE FALSE END is_saved "
			s[1] = " LEFT JOIN (SELECT DISTINCT job_id FROM saved_jobs WHERE user_id =" + userId + ") sj ON sj.job_id = jobs.id "
			s[2] = " ,sj.job_id "
		}

	}
	var jobTitle string
	row.Scan(&jobTitle)

	rows, err := db.DB.Query(ctx.Context(), "SELECT jobs.id, title, companies.name, remote_available, salary_from, salary_to "+s[0]+" ,array_agg(distinct tech_skills.name) AS skills"+" FROM jobs JOIN companies ON companies.id = company_id "+s[1]+`LEFT JOIN job_tech_skills ON job_tech_skills.job_id = jobs.id
	LEFT JOIN tech_skills ON job_tech_skills.tech_skill_id = tech_skills.id WHERE levenshtein(title, $1) <= 4 AND jobs.id != $2 GROUP BY jobs.id, companies.name `+s[2]+"LIMIT 8", jobTitle, jobId)
	if err != nil {
		slog.Error("Query error: ", "err", err)
	}
	var result RelatedJob
	var results []RelatedJob

	result.Salary = &RelatedJobSalary{}

	if s[0] != "" {
		for rows.Next() {
			err = rows.Scan(&result.Id, &result.Title, &result.CompanyName, &result.RemoteAvailable, &result.Salary.From, &result.Salary.To, &result.IsSaved, &result.Skills)
			if err != nil {
				slog.Error("Scan error", "err", err)
			}
			results = append(results, result)
		}
	} else {
		for rows.Next() {
			err = rows.Scan(&result.Id, &result.Title, &result.CompanyName, &result.RemoteAvailable, &result.Salary.From, &result.Salary.To, &result.Skills)
			if err != nil {
				slog.Error("Scan error", "err", err)
			}
			results = append(results, result)
		}
	}

	return ctx.JSON(RelatedJobsResponse{
		Results: results,
	})
}

type RelatedJobsResponse struct {
	Results []RelatedJob `json:"results"`
}
