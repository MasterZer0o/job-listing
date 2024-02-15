package seeders

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"main/db"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v5"
)

type JobContent struct {
	Content  string   `json:"content"`
	Sections []string `json:"sections"`
}

type JobContentsRow struct {
	Content  string
	Sections []string
	JobId    int
}

type JobData struct {
	Id       int
	ExpLevel string
}

func SeedContents() {
	jobsData := getJobsData()
	jobsDataLen := len(jobsData)
	slog.Info("Starting seeding job contents...", "rows to create", jobsDataLen)

	const VALUES_PER_WORKER = 500
	workerCount := 1
	if jobsDataLen > VALUES_PER_WORKER {
		if jobsDataLen%500 != 0 {
			workerCount = jobsDataLen/(VALUES_PER_WORKER) + 1
		} else {
			workerCount = jobsDataLen / (VALUES_PER_WORKER)
		}
	}

	isDividableByValuesPerWorker := jobsDataLen%500 == 0
	var jobIdsShare [][]JobData
	startIdx, endIdx := 0, VALUES_PER_WORKER

	// allocating worker share
	if isDividableByValuesPerWorker {
		for i := 0; i < workerCount; i++ {
			jobIdsShare = append(jobIdsShare, jobsData[startIdx:endIdx])
			startIdx = endIdx
			endIdx += VALUES_PER_WORKER
		}
	} else {
		if workerCount-1 == 0 {
			jobIdsShare = append(jobIdsShare, jobsData[0:jobsDataLen])
		} else {
			for i := 0; i < workerCount-1; i++ {
				jobIdsShare = append(jobIdsShare, jobsData[startIdx:endIdx])
				startIdx = endIdx
				endIdx += VALUES_PER_WORKER
			}
			jobIdsShare = append(jobIdsShare, jobsData[startIdx:])
		}

	}
	generatedContents := generateContents(jobIdsShare, jobsDataLen, workerCount)
	insertContents(generatedContents)
}

func generateContents(jobIdsShare [][]JobData, jobIdsLen int, workerCount int) []JobContentsRow {
	slog.Info("Generating contents", "workers count", workerCount)

	var activeWorkers atomic.Uint32
	activeWorkers.Store(uint32(workerCount))
	wg := sync.WaitGroup{}
	contents := getJobContents()

	valuesC := make(chan JobContentsRow, jobIdsLen)
	for i := 0; i < len(jobIdsShare); i++ {
		wg.Add(1)

		go func(share []JobData, wg *sync.WaitGroup, vc chan JobContentsRow, counter *atomic.Uint32) {
			for j := 0; j < len(share); j++ {
				var idx int
				switch share[j].ExpLevel {
				case "Junior":
					idx = 0
				case "Mid":
					idx = 1
				default:
					idx = 2
				}

				valuesC <- JobContentsRow{
					Content:  contents[idx].Content,
					Sections: contents[idx].Sections,
					JobId:    share[j].Id,
				}
			}

			wg.Done()
			counter.Add(^uint32(0))
		}(jobIdsShare[i], &wg, valuesC, &activeWorkers)
	}
	var values []JobContentsRow
	wg.Add(1)
	go func(activeWorkers *atomic.Uint32) {
		for range valuesC {
			v := <-valuesC
			values = append(values, v)
			if activeWorkers.Load() == 0 && len(valuesC) == 0 {
				close(valuesC)
				wg.Done()
			}
		}
	}(&activeWorkers)
	wg.Wait()
	slog.Info("Finished generating contents", "rows", len(values))
	return values
}

func getJobContents() []JobContent {
	fileName := "seed_data.json"
	v, err := os.ReadFile("db/" + fileName)
	if err != nil {
		slog.Error("Failed to read job contents file", "file name", fileName, "err", err)
	}

	contents := []JobContent{}
	err = json.Unmarshal(v, &contents)
	if err != nil {
		slog.Error("Error marshaling JSON into struct", "err", err)
	}
	return contents
}

func insertContents(input []JobContentsRow) {
	start := time.Now()
	rowsCopied, err := db.DB.CopyFrom(context.Background(), pgx.Identifier{"jobs_content"}, []string{
		"content", "sections", "job_id",
	}, pgx.CopyFromSlice(len(input), func(i int) ([]interface{}, error) {
		return []interface{}{
			input[i].Content, input[i].Sections, input[i].JobId,
		}, nil
	}))

	if err != nil {
		slog.Error("Error inserting data to database", "err", err)
		return
	}

	slog.Info(fmt.Sprintf("Finished inserting jobs content, rows inserted %d/%d", rowsCopied, len(input)),
		"time", time.Since(start))
}

func getJobsData() []JobData {
	rows, err := db.DB.Query(context.Background(), "SELECT jobs.id, SUBSTRING(jobs.title,0, POSITION(' ' in jobs.title)) FROM jobs")
	if err != nil {
		slog.Error("Error reading JobData", "err", err)
	}
	v := JobData{}
	var vs []JobData
	for rows.Next() {
		rows.Scan(&v.Id, &v.ExpLevel)
		vs = append(vs, v)
	}
	rows.Close()
	return vs
}
