package seeders

import (
	"context"
	"fmt"
	"main/db"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
)

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

func SeedJobs(rowCount uint64) {
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

	rowsCopied, err := db.DB.CopyFrom(context.Background(), pgx.Identifier{"jobs"},
		[]string{"location", "remote_available", "company_id", "level", "title", "salary_from", "salary_to",
			"currency_id"}, pgx.CopyFromSlice(len(rowSources), func(i int) ([]interface{}, error) {
			return []interface{}{
				&rowSources[i].Location, &rowSources[i].RemoteAvailable, &rowSources[i].CompanyId, &rowSources[i].Level, &rowSources[i].Title, &rowSources[i].SalaryFrom, &rowSources[i].SalaryTo, &rowSources[i].CurrencyId}, nil
		}))

	if err != nil {
		fmt.Println("Error from inserting rows", err)
	}
	fmt.Printf("Seeding %d jobs took: %s.\nJob rows inserted: %d/%d\n", rowCount, time.Since(start), rowsCopied, rowCount)
	SeedSkills()

}

func generateJobs(workers uint64, rowSplit []uint64, jobC chan<- SeededJob) {
	randomLocationsCount := 20
	locations := []string{}
	for i := 0; i < randomLocationsCount; i++ {
		locations = append(locations, "Location"+strconv.Itoa(rand.Intn(50)))
	}
	levels := [3]string{"Junior", "Mid", "Senior"}

	rows, _ := db.DB.Query(context.Background(), "SELECT id FROM companies LIMIT 50")
	var companyIds []uint16

	for rows.Next() {
		var id uint16
		rows.Scan(&id)
		companyIds = append(companyIds, id)
	}
	rows.Close()

	rows, _ = db.DB.Query(context.Background(), "SELECT id FROM currencies LIMIT 50")
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
			fmt.Printf("Worker #%d finished generating data. [%s]\n", identifier+1, time.Since(start))
		}(jobC, &v, i)

	}
	wg.Wait()
	close(jobC)
}
