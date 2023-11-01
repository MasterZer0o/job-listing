package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
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
	countInt, _ := strconv.ParseUint(count, 10, 64)

	switch seedTarget {
	case "jobs":
		seedJobs(countInt)
	case "users":

	}
	return ctx.JSON(map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Successfully inserted %d rows into %s", countInt, seedTarget),
	})
}

type SeededJob struct {
	Location        string
	RemoteAvailable bool
	CompanyId       uint16
	Level           string
	Title           string
	SalaryFrom      uint16
	SalaryTo        uint16
	CurrencyId      uint16
}

func seedJobs(rowCount uint64) {
	start := time.Now()

	rowSources := make([]SeededJob, 0, rowCount)
	jobC := make(chan SeededJob, rowCount/2)

	workers := uint64(3)
	rowSplit := make([]uint64, 0, workers)

	// if workers <= rowCount case not covered->tofix
	if rowCount >= workers {
		dividedInt := uint64(rowCount / workers)
		if rowCount%workers == 0 {
			for i := uint64(0); i < workers; i++ {
				rowSplit = append(rowSplit, dividedInt)
			}
		} else {
			for i := uint64(0); i < workers-1; i++ {
				rowSplit = append(rowSplit, dividedInt)
			}
			rowSplit = append(rowSplit, dividedInt+rowCount%workers)
		}
	}

	go generateJobs(workers, rowSplit, jobC)

	for entry := range jobC {
		rowSources = append(rowSources, entry)
	}

	rowsCopied, err := DB.CopyFrom(context.Background(), pgx.Identifier{"jobs"},
		[]string{"location", "remote_available", "company_id", "level", "title", "salary_from", "salary_to",
			"currency_id"}, pgx.CopyFromSlice(len(rowSources), func(i int) ([]interface{}, error) {
			return []interface{}{
				&rowSources[i].Location, &rowSources[i].RemoteAvailable, &rowSources[i].CompanyId, &rowSources[i].Level, &rowSources[i].Title, &rowSources[i].SalaryFrom, &rowSources[i].SalaryTo, &rowSources[i].CurrencyId}, nil
		}))

	if err != nil {
		fmt.Println("Error from copying job rows", err)
	}
	elapsed := time.Since(start)
	fmt.Printf("Seeding %d jobs took: %s.\nJob rows inserted: %d/%d\n", rowCount, elapsed, rowsCopied, rowCount)
	seedSkills(rowCount)

}

func generateJobs(workers uint64, rowSplit []uint64, jobC chan<- SeededJob) {
	randomLocationsCount := 20
	locations := []string{}
	for i := 0; i < randomLocationsCount; i++ {
		locations = append(locations, "Location"+strconv.Itoa(rand.Intn(50)))
	}
	levels := [3]string{"Junior", "Mid", "Senior"}

	rows, _ := DB.Query(context.Background(), "SELECT id FROM companies LIMIT 20")
	var companyIds []uint16

	for rows.Next() {
		var id uint16
		rows.Scan(&id)
		companyIds = append(companyIds, id)
	}
	rows.Close()

	rows, _ = DB.Query(context.Background(), "SELECT id FROM currencies LIMIT 20")
	var currencyIds []uint16

	for rows.Next() {
		var id uint16
		rows.Scan(&id)
		currencyIds = append(currencyIds, id)
	}
	rows.Close()

	wg := sync.WaitGroup{}

	for i, v := range rowSplit {
		wg.Add(1)
		go func(c chan<- SeededJob, rowCount *uint64, identifier int) {
			defer wg.Done()
			start := time.Now()

			for j := 0; j < int(*rowCount); j++ {
				remoteAvailable := rand.Intn(2)
				level := levels[rand.Intn(3)]
				title := level + " " + "Title" + strconv.Itoa(rand.Intn(100))
				randomLocation := locations[rand.Intn(randomLocationsCount)]
				salaryFrom := rand.Intn(6000-3200) + 3200
				salaryTo := rand.Intn(15500-salaryFrom) + salaryFrom

				jobC <- SeededJob{
					Location:        randomLocation,
					RemoteAvailable: remoteAvailable == 0,
					CompanyId:       companyIds[rand.Intn(len(companyIds))],
					Level:           level,
					Title:           title,
					SalaryFrom:      uint16(salaryFrom),
					SalaryTo:        uint16(salaryTo),
					CurrencyId:      currencyIds[rand.Intn(len(currencyIds))],
				}
			}
			elapsed := time.Since(start)
			fmt.Printf("Worker #%d finished generating data. [%s]\n", identifier+1, elapsed)
		}(jobC, &v, i)

	}
	wg.Wait()
	close(jobC)
}

type Skill struct {
	name  string
	jobId jobId
}
type jobId int

func seedSkills(jobCount uint64) {
	const MAX_NUMBER_OF_SKILLS = 6

	fmt.Println("\nSeeding job skills...")
	start := time.Now()
	rows, err := DB.Query(context.Background(), "SELECT id FROM jobs ORDER BY id ASC")

	if err != nil {
		log.Fatal("Error reading job ids: ", err)
	}

	var jobIds []jobId
	var id jobId
	for rows.Next() {
		rows.Scan(&id)
		jobIds = append(jobIds, id)
	}
	rows.Close()

	var workers int = 1

	jobIdsCount := len(jobIds)
	var divisor int

	if jobIdsCount > 500 {
		divisor = 500
	} else {
		divisor = jobIdsCount / 2
	}

	// len(jobIds) is less than 500
	if jobIdsCount < divisor {
		if jobIdsCount%2 == 0 {
			workers = 2
		}
	} else {
		if jobIdsCount%divisor == 0 {
			workers = jobIdsCount / divisor
		} else {
			workers = jobIdsCount/divisor + 1
		}
	}

	fmt.Println("Workers: ", workers)
	dataSplit := make([][]jobId, 0, workers)

	indexStart, indexEnd := 0, divisor

	if jobIdsCount%divisor == 0 {
		for i := 0; i < workers; i++ {
			dataSplit = append(dataSplit, jobIds[indexStart:indexEnd])

			indexStart = indexEnd
			indexEnd += divisor

		}
	} else {
		for i := 0; i < workers-1; i++ {
			dataSplit = append(dataSplit, jobIds[indexStart:indexEnd])

			indexStart = indexEnd
			indexEnd += divisor
		}
		dataSplit = append(dataSplit, jobIds[indexEnd:jobIdsCount])

	}

	jobSkillsRows := []Skill{}
	var wg sync.WaitGroup
	for i, v := range dataSplit {
		wg.Add(1)
		go func(jobIds []jobId, id int) {
			defer wg.Done()
			start := time.Now()
			for _, jobId := range jobIds {
				numbOfSkills := rand.Intn(MAX_NUMBER_OF_SKILLS) + 1

				for j := 0; j < numbOfSkills; j++ {
					jobSkillsRows = append(jobSkillsRows, Skill{
						name:  "Skill" + strconv.Itoa(j+1),
						jobId: jobId,
					})

				}
			}
			elapsed := time.Since(start)
			fmt.Printf("Worker #%d finished generating skills for %d jobs in %s.\n", id, len(jobIds), elapsed)
		}(v, i+1)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Generating skills for %d took %s.\n", jobCount, elapsed)
	fmt.Println("Inserting to database...")

	start = time.Now()
	rowsCopied, err := DB.CopyFrom(context.Background(), pgx.Identifier{"job_skills"},
		[]string{"name", "job_id"}, pgx.CopyFromSlice(len(jobSkillsRows), func(i int) ([]any, error) {
			return []any{&jobSkillsRows[i].name, &jobSkillsRows[i].jobId}, nil
		}))

	if err != nil {
		fmt.Println("Error while inserting skills to database: ", err)
		return
	}

	elapsed = time.Since(start)
	fmt.Printf("Inserted %d/%d job skills in %s.", rowsCopied, len(jobSkillsRows), elapsed)

}
