package seeders

import (
	"context"
	"fmt"
	"log/slog"
	"main/db"
	"math/rand"
	"sync"
	"time"
)

type JobId = string
type Value struct {
	jobId    JobId
	filterId string
}

type AllFiltersPossibleValues struct {
	empTypes    PossibleValues
	typesOfWork PossibleValues
	workModes   PossibleValues
	techSkills  PossibleValues
	expLevels   PossibleValues
}

type ValueInputs struct {
	empTypes    chan Value
	typesOfWork chan Value
	workModes   chan Value
	techSkills  chan Value
	expLevels   chan Value
}

func SeedFilters() {
	empTypes := GetPossibleValues("employment_types")
	typesOfWork := GetPossibleValues("types_of_work")
	workModes := GetPossibleValues("work_modes")
	techSkills := GetPossibleValues("tech_skills")
	expLevels := GetPossibleValues("experience_levels")
	allPosValues := AllFiltersPossibleValues{
		empTypes,
		typesOfWork,
		workModes,
		techSkills,
		expLevels,
	}

	jobIds := []JobId{}
	var v JobId
	rows, _ := db.DB.Query(context.Background(), "SELECT id from jobs")
	for rows.Next() {
		rows.Scan(&v)
		jobIds = append(jobIds, v)
	}
	rows.Close()
	jobIdsCount := len(jobIds)

	wg := sync.WaitGroup{}
	const VALUES_PER_WORKER = 500
	isDividableByValuesPerWorker := jobIdsCount%500 == 0
	// share := jobIdsCount / VALUES_PER_WORKER
	workers := 1
	if jobIdsCount > VALUES_PER_WORKER {
		if isDividableByValuesPerWorker {
			workers = jobIdsCount / VALUES_PER_WORKER
		} else {
			workers = jobIdsCount/VALUES_PER_WORKER + 1
		}
	}
	workerShare := make([][]JobId, 0, workers)
	startIdx := 0
	endIdx := VALUES_PER_WORKER

	if isDividableByValuesPerWorker {
		for i := 0; i < workers; i++ {
			workerShare = append(workerShare, jobIds[startIdx:endIdx])
			startIdx = endIdx
			endIdx += VALUES_PER_WORKER
		}
	} else {
		if workers-1 == 0 {
			workerShare = append(workerShare, jobIds[0:jobIdsCount])
		} else {
			for i := 0; i < workers-1; i++ {
				workerShare = append(workerShare, jobIds[startIdx:endIdx])
				startIdx = endIdx
				endIdx += VALUES_PER_WORKER
			}
			workerShare = append(workerShare, jobIds[startIdx:])
		}

	}
	valueInputs := ValueInputs{
		empTypes:    make(chan Value, 2*workers),
		typesOfWork: make(chan Value, 2*workers),
		workModes:   make(chan Value, 2*workers),
		techSkills:  make(chan Value, 2*workers),
		expLevels:   make(chan Value, 2*workers),
	}
	timeStart := time.Now()
	for idx, share := range workerShare {
		wg.Add(1)
		go jobFilterGenerator(&valueInputs, &allPosValues, &wg, share, idx+1)

	}

	wg.Wait()
	slog.Info(fmt.Sprintf("%d Finished generating data", workers), "time spent", time.Since(timeStart))

}

// TODO: drain values from channels
func jobFilterGenerator(inputs *ValueInputs, possVals *AllFiltersPossibleValues, wg *sync.WaitGroup, share []JobId, workerId int) {
	slog.Info("start", "workerId", workerId)
	timeStart := time.Now()
	for _, jobId := range share {
		for i := 0; i < rand.Intn(6)+1; i++ {
			inputs.techSkills <- Value{
				jobId:    jobId,
				filterId: possVals.techSkills.values[rand.Intn(possVals.techSkills.len)],
			}
		}

		for i := 0; i < rand.Intn(5)+1; i++ {
			inputs.empTypes <- Value{
				jobId:    jobId,
				filterId: possVals.empTypes.values[rand.Intn(possVals.empTypes.len)],
			}
		}

		for i := 0; i < rand.Intn(3)+1; i++ {
			inputs.expLevels <- Value{
				jobId:    jobId,
				filterId: possVals.expLevels.values[rand.Intn(possVals.expLevels.len)],
			}
		}

		for i := 0; i < rand.Intn(3)+1; i++ {
			inputs.typesOfWork <- Value{
				jobId:    jobId,
				filterId: possVals.typesOfWork.values[rand.Intn(possVals.typesOfWork.len)],
			}
		}

		for i := 0; i < rand.Intn(3)+1; i++ {
			inputs.workModes <- Value{
				jobId:    jobId,
				filterId: possVals.workModes.values[rand.Intn(possVals.workModes.len)],
			}
		}
	}

	wg.Done()
	slog.Info(fmt.Sprintf("Worker %d finished generating data.", workerId), "time spent", time.Since(timeStart))
}

func insertFilters() {

}
