package seeders

import (
	"context"
	"fmt"
	"log/slog"
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
	CompanyId       string
	Title           string
	SalaryFrom      uint16
	SalaryTo        uint16
	CurrencyId      string
	Level           string
}

func SeedJobs(rowCount int, idReset bool) {
	start := time.Now()

	rowSources := make([]SeededJob, 0, rowCount)
	jobC := make(chan SeededJob, rowCount/2)

	workers := 3
	rowSplit := make([]int, 0, workers)

	if rowCount >= workers {
		dividedInt := rowCount / workers
		if rowCount%workers == 0 {
			for i := 0; i < workers; i++ {
				rowSplit = append(rowSplit, dividedInt)
			}
		} else {
			for i := 0; i < workers-1; i++ {
				rowSplit = append(rowSplit, dividedInt)
			}
			rowSplit = append(rowSplit, dividedInt+rowCount%workers)
		}
	}
	if idReset {
		_, err := db.DB.Exec(context.Background(), "ALTER SEQUENCE jobs_id_seq RESTART WITH 1")
		if err != nil {
			slog.Error("Error resetting auto-incrementing id", "err", err)
		} else {
			slog.Info("Id has been reset.")
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
				rowSources[i].Location, rowSources[i].RemoteAvailable, rowSources[i].CompanyId, rowSources[i].Level, rowSources[i].Title, rowSources[i].SalaryFrom, rowSources[i].SalaryTo, rowSources[i].CurrencyId}, nil
		}))

	if err != nil {
		slog.Error("Error from inserting rows", "err", err)
		return
	}

	slog.Info(fmt.Sprintf("Seeding %d jobs took: %s.\nJob rows inserted: %d/%d\n",
		rowCount, time.Since(start), rowsCopied, rowCount))

	SeedFilters()
}

func generateJobs(workers int, rowSplit []int, jobC chan<- SeededJob) {
	randomLocationsCount := 20
	locations := []string{}
	for i := 0; i < randomLocationsCount; i++ {
		locations = append(locations, "Location"+strconv.Itoa(rand.Intn(50)))
	}
	companyIds := GetPossibleValues("companies")
	currencyIds := GetPossibleValues("currencies")

	wg := sync.WaitGroup{}
	for i, v := range rowSplit {
		wg.Add(1)
		go func(c chan<- SeededJob, rowCount *int, workerId int) {
			defer wg.Done()
			start := time.Now()
			rc := *rowCount
			var titleLevel string
			randLevel := rand.Intn(3) + 1
			switch randLevel {
			case 1:
				titleLevel = "Junior"
			case 2:
				titleLevel = "Mid"
			case 3:
				titleLevel = "Senior"
			}

			for j := 0; j < rc; j++ {
				remoteAvailable := rand.Intn(2) == 0
				title := titleLevel + " " + "Title" + strconv.Itoa(rand.Intn(100))
				randomLocation := locations[rand.Intn(randomLocationsCount)]
				salaryFrom := rand.Intn(6000-3200) + 3200
				salaryTo := rand.Intn(15500-salaryFrom) + salaryFrom
				randomCompanyId := companyIds.values[rand.Intn(companyIds.len)]
				randomCurrencyId := currencyIds.values[rand.Intn(currencyIds.len)]

				jobC <- SeededJob{
					Location:        randomLocation,
					RemoteAvailable: remoteAvailable,
					CompanyId:       fmt.Sprint(randomCompanyId),
					Title:           title,
					SalaryFrom:      uint16(salaryFrom),
					SalaryTo:        uint16(salaryTo),
					CurrencyId:      fmt.Sprint(randomCurrencyId),
					Level:           titleLevel,
				}
			}
			slog.Info(fmt.Sprintf("Worker #%d finished generating data. [%s]\n", workerId+1, time.Since(start)))
		}(jobC, &v, i)

	}
	wg.Wait()
	close(jobC)
}

type PossibleValues struct {
	values []string
	len    int
}

func GetPossibleValues(tableName string) PossibleValues {
	rows, err := db.DB.Query(context.Background(), fmt.Sprintf("SELECT id from %s LIMIT 50", tableName))

	if err != nil {
		slog.Error("Fetching possible values failed", slog.String("table name", tableName))
		return PossibleValues{}
	}

	var value string
	var values []string
	for rows.Next() {
		rows.Scan(&value)
		values = append(values, value)
	}
	rows.Close()

	return PossibleValues{
		values: values,
		len:    len(values),
	}
}
