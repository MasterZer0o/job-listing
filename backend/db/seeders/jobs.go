package seeders

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log/slog"
	"main/db"
	"math/rand"

	"strconv"
	"sync"
	"time"

	"github.com/fogleman/gg"

	"github.com/jackc/pgx/v5"
)

type SeededJob struct {
	Location        string
	RemoteAvailable bool
	CompanyId       string
	Title           string
	SalaryMin       int
	SalaryMax       int
	CurrencyId      string
	Level           string
}

type SeededCompany struct {
	Name  string
	Image string
}

type Salary struct {
	Min, Max int
}

func SeedJobs(jobCount int, idReset bool) {
	rowSources := make([]SeededJob, 0, jobCount)
	jobC := make(chan SeededJob, jobCount/2)

	generatedCompanies := generateCompanies(jobCount)
	insertCompanies(generatedCompanies, len(generatedCompanies))
	companiesIds := make(chan string, jobCount)
	companiesIdsArr := getCompaniesId(companiesIds)
	companiesIdsLen := len(companiesIdsArr)

	// fill companiesIds with duplicates to match jobs count
	for companiesIdsLen < jobCount {
		companiesIds <- companiesIdsArr[rand.Intn(companiesIdsLen)]
	}

	jobContents := getJobContents()

	workers := 3
	rowSplit := make([]int, 0, workers)

	if jobCount >= workers {
		dividedInt := jobCount / workers
		if jobCount%workers == 0 {
			for i := 0; i < workers; i++ {
				rowSplit = append(rowSplit, dividedInt)
			}
		} else {
			for i := 0; i < workers-1; i++ {
				rowSplit = append(rowSplit, dividedInt)
			}
			rowSplit = append(rowSplit, dividedInt+jobCount%workers)
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
	go generateJobs(rowSplit, jobC, companiesIds, jobContents)

	for entry := range jobC {
		rowSources = append(rowSources, entry)
	}

	insertJobs(rowSources, jobCount)
	jobIds := getAllJobIds()
	go SeedFilters(jobIds)
	SeedContents()
}

func generateJobs(rowSplit []int, jobC chan<- SeededJob, companiesIds chan string, jobContents []JobContent) {

	currencyIds := GetPossibleValues("currencies")

	wg := sync.WaitGroup{}
	for i, v := range rowSplit {
		wg.Add(1)
		go func(c chan<- SeededJob, rowCount *int, workerId int) {
			defer wg.Done()
			start := time.Now()
			rc := *rowCount

			for j := 0; j < rc; j++ {
				expLevel := generateExpLevel()
				remoteAvailable := rand.Intn(2) == 0
				title := generateTitle(expLevel)
				randomLocation := generateLocation()
				salary := generateSalary(expLevel)
				randomCompanyId := <-companiesIds
				randomCurrencyId := currencyIds.values[rand.Intn(currencyIds.len)]

				jobC <- SeededJob{
					Location:        randomLocation,
					RemoteAvailable: remoteAvailable,
					CompanyId:       randomCompanyId,
					Title:           title,
					SalaryMin:       salary.Min,
					SalaryMax:       salary.Max,
					CurrencyId:      fmt.Sprint(randomCurrencyId),
					Level:           expLevel,
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

func calculateWorkerCount() {}
func generateImage(companyName string) string {
	const (
		width    = 160
		height   = 160
		fontSize = 20
		margin   = 10
	)

	// Create a new image
	image.NewRGBA(image.Rect(0, 0, width, height))
	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Set font properties
	dc.LoadFontFace("", fontSize)

	// Draw the text with word wrap
	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped(companyName, margin, margin, 0, 0, float64(width-2*margin), 1.5, gg.AlignCenter)

	// Encode the image to PNG format
	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		slog.Error("Error encoding image", "err", err)
		return ""
	}

	// Encode the PNG image to base64
	encodedString := base64.StdEncoding.EncodeToString(buf.Bytes())

	return "data:image/png;base64," + encodedString
}
func generateCompanies(jobCount int) []SeededCompany {
	companiesCount := jobCount/rand.Intn(2) + 2
	var results []SeededCompany
	for i := 0; i < companiesCount; i++ {
		name := "Best comp " + strconv.Itoa(i)
		results = append(results, SeededCompany{
			Name:  name,
			Image: generateImage(name),
		})
	}
	slog.Info("Generating companies finished.")
	return results
}
func generateSalary(expLeveL string) Salary {
	const thresholds = 3
	var min, max [thresholds]int
	switch expLeveL {
	case "Junior":
		min, max = [thresholds]int{3500, 4000, 4500}, [thresholds]int{6000, 6500, 7000}
	case "Mid":
		min, max = [thresholds]int{5500, 6000, 7000}, [thresholds]int{8000, 9500, 10000}
	case "Senior":
		min, max = [thresholds]int{10500, 11000, 12000}, [thresholds]int{12000, 14500, 16000}
	}

	randIdx := rand.Intn(thresholds) + 1
	return Salary{
		Min: min[randIdx],
		Max: max[randIdx],
	}
}
func generateExpLevel() string {
	var expLevel string
	randLevel := rand.Intn(3) + 1
	switch randLevel {
	case 1:
		expLevel = "Junior"
	case 2:
		expLevel = "Mid"
	case 3:
		expLevel = "Senior"
	}
	return expLevel
}
func generateLocation() string {
	locations := []string{
		"USA, New York",
		"UK, London",
		"France, Paris",
		"Japan, Tokyo",
		"China, Beijing",
		"Germany, Berlin",
		"Italy, Rome",
		"Canada, Toronto",
		"Australia, Sydney",
		"Brazil, Rio de Janeiro",
		"India, Mumbai",
		"South Africa, Johannesburg",
		"Russia, Moscow",
		"Spain, Madrid",
		"South Korea, Seoul",
		"Mexico, Mexico City",
		"Argentina, Buenos Aires",
		"Netherlands, Amsterdam",
		"Switzerland, Zurich",
		"Sweden, Stockholm",
		"Belgium, Brussels",
		"Austria, Vienna",
		"Norway, Oslo",
		"Denmark, Copenhagen",
		"Greece, Athens",
		"Portugal, Lisbon",
		"Turkey, Istanbul",
		"Thailand, Bangkok",
		"Egypt, Cairo",
		"Indonesia, Jakarta",
		"Poland, Warsaw",
		"Saudi Arabia, Riyadh",
		"United Arab Emirates, Dubai",
		"Singapore, Singapore",
		"Malaysia, Kuala Lumpur",
		"Vietnam, Ho Chi Minh City",
		"Philippines, Manila",
		"Israel, Tel Aviv",
		"Ireland, Dublin",
		"Finland, Helsinki",
	}

	return locations[rand.Intn(len(locations))]
}
func generateDescription() {
	// TODO:
}
func generateTitle(expLevel string) string {
	// TODO:
	return expLevel + " " + "Title" + strconv.Itoa(rand.Intn(100))
}

func insertJobs(rowSources []SeededJob, rowCount int) {
	timeStart := time.Now()

	rowsCopied, err := db.DB.CopyFrom(context.Background(), pgx.Identifier{"jobs"},
		[]string{"location", "remote_available", "company_id", "level", "title", "salary_from", "salary_to",
			"currency_id"}, pgx.CopyFromSlice(len(rowSources), func(i int) ([]interface{}, error) {
			return []interface{}{
				rowSources[i].Location, rowSources[i].RemoteAvailable, rowSources[i].CompanyId, rowSources[i].Level, rowSources[i].Title, rowSources[i].SalaryMin, rowSources[i].SalaryMax, rowSources[i].CurrencyId}, nil
		}))

	if err != nil {
		slog.Error("Error from inserting rows [jobs]", "err", err)
		return
	}

	slog.Info(fmt.Sprintf("Finished seeding %d jobs [%s].\nRows inserted: %d/%d\n",
		rowCount, time.Since(timeStart), rowsCopied, rowCount))
}
func insertCompanies(rowSources []SeededCompany, rowCount int) {
	timeStart := time.Now()

	rowsCopied, err := db.DB.CopyFrom(context.Background(), pgx.Identifier{"companies"},
		[]string{"name", "image"}, pgx.CopyFromSlice(len(rowSources), func(i int) ([]interface{}, error) {
			return []interface{}{rowSources[i].Name, rowSources[i].Image}, nil
		}))

	if err != nil {
		slog.Error("Error from inserting rows [companies]", "err", err)
		return
	}

	slog.Info(fmt.Sprintf("Finished seeding %d companies [%s].\nRows inserted: %d/%d\n",
		rowCount, time.Since(timeStart), rowsCopied, rowCount))
}

func getCompaniesId(idsC chan string) []string {
	rows, err := db.DB.Query(context.Background(), "SELECT id from companies")
	if err != nil {
		slog.Error("Error getting companies id", "err", err)
	}
	var values []string
	var value string
	for rows.Next() {
		rows.Scan(&value)
		idsC <- value
		values = append(values, value)
	}
	rows.Close()
	return values
}

type JobId = string

func getAllJobIds() []JobId {
	var v JobId
	var jobIds []JobId
	rows, _ := db.DB.Query(context.Background(), "SELECT id from jobs")
	for rows.Next() {
		rows.Scan(&v)
		jobIds = append(jobIds, v)
	}
	rows.Close()
	return jobIds
}
