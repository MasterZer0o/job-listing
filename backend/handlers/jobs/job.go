package handlers

import (
	"log/slog"
	"main/db"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CompanyData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type JobData struct {
	Title           string      `json:"title"`
	SalaryFrom      int         `json:"salaryFrom"`
	SalaryTo        int         `json:"salaryTo"`
	Level           string      `json:"level"`
	Company         CompanyData `json:"company"`
	RemoteAvailable bool        `json:"remote"`
	CreatedAt       time.Time   `json:"createdAt"`
	Skills          []string    `json:"skills"`
	IsSaved         bool        `json:"isSaved"`
}

func GetJob(ctx *fiber.Ctx) error {
	jobId := ctx.Params("id")

	t := time.Now()
	row := db.DB.QueryRow(ctx.Context(), `SELECT title, salary_from, salary_to,level, remote_available, companies.name, companies.image, companies.id, jobs.created_at, array_agg(tech_skills.name), (SELECT EXISTS (SELECT saved_jobs.job_id FROM sessions
		JOIN users ON users.id = sessions.user_id
		JOIN saved_jobs ON saved_jobs.user_id = users.id
		WHERE sessions.id=$1 AND saved_jobs.job_id=$2
		LIMIT 1)) FROM jobs
	JOIN companies on jobs.company_id = companies.id
  JOIN job_tech_skills ON job_tech_skills.job_id = jobs.id
  JOIN tech_skills ON tech_skills.id = job_tech_skills.tech_skill_id
	WHERE jobs.id = $2
  GROUP BY jobs.title, jobs.salary_from, jobs.salary_to, jobs.level, jobs.remote_available,companies.name,companies.image,companies.id, jobs.created_at`, ctx.Cookies("session"), jobId)

	slog.Info("Time taken to query job", "job id", jobId, "time", time.Since(t))

	var data JobData
	err := row.Scan(&data.Title, &data.SalaryFrom, &data.SalaryTo, &data.Level, &data.RemoteAvailable, &data.Company.Name, &data.Company.Image, &data.Company.Id, &data.CreatedAt, &data.Skills, &data.IsSaved)

	if err != nil {
		slog.Error("Scan Error", "err", err)
	}

	return ctx.JSON(map[string]interface{}{
		"data": data,
	})
}
