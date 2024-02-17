package seeders

import (
	"context"
	"fmt"
	"log/slog"
	"main/db"
	"math/rand"
	"reflect"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
)

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
	EmpTypes    chan Value
	TypesOfWork chan Value
	WorkModes   chan Value
	TechSkills  chan Value
	ExpLevels   chan Value
}

func SeedFilters(jobIds []JobId) {
	slog.Info("Starting seeding job filters")
	fullStartTime := time.Now()

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

	jobIdsCount := len(jobIds)

	wg := sync.WaitGroup{}
	const VALUES_PER_WORKER = 500
	isDividableByValuesPerWorker := jobIdsCount%500 == 0
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
		EmpTypes:    make(chan Value, 2*workers),
		TypesOfWork: make(chan Value, 2*workers),
		WorkModes:   make(chan Value, 2*workers),
		TechSkills:  make(chan Value, 2*workers),
		ExpLevels:   make(chan Value, 2*workers),
	}

	for idx, share := range workerShare {
		wg.Add(1)
		go jobFilterGenerator(&valueInputs, &allPosValues, &wg, share, idx+1)
	}

	outputs := Outputs{
		"EmpTypes":    {},
		"TypesOfWork": {},
		"WorkModes":   {},
		"TechSkills":  {},
		"ExpLevels":   {},
	}

	nwg := sync.WaitGroup{}
	go drainInputs(&valueInputs, outputs, &nwg)
	wg.Wait()

	close(valueInputs.EmpTypes)
	close(valueInputs.ExpLevels)
	close(valueInputs.WorkModes)
	close(valueInputs.TypesOfWork)
	close(valueInputs.TechSkills)

	nwg.Wait()

	insertingWg := sync.WaitGroup{}

	go insertFilters("job_employment_types", "emp_type_id", *outputs["EmpTypes"], &insertingWg)
	go insertFilters("job_experience_levels", "exp_level_id", *outputs["ExpLevels"], &insertingWg)
	go insertFilters("job_tech_skills", "tech_skill_id", *outputs["TechSkills"], &insertingWg)
	go insertFilters("job_work_modes", "work_mode_id", *outputs["WorkModes"], &insertingWg)
	go insertFilters("job_types_of_work", "type_of_work_id", *outputs["TypesOfWork"], &insertingWg)
	insertingWg.Wait()

	slog.Info("Seeding filters finished.", "time taken", time.Since(fullStartTime))
}

type Outputs = map[string]*[]Value

func drainInputs(inputs *ValueInputs, outputs Outputs, parentWg *sync.WaitGroup) {
	wg := sync.WaitGroup{}
	parentWg.Add(1)
	structV := reflect.ValueOf(*inputs)

	for i := 0; i < structV.NumField(); i++ {
		wg.Add(1)

		inputChan := structV.Field(i).Interface().(chan Value)
		outputSlc := outputs[structV.Type().Field(i).Name]

		go func(input chan Value, output *[]Value) {
			for v := range input {
				*output = append(*output, v)
			}
			wg.Done()

		}(inputChan, outputSlc)
	}
	wg.Wait()
	parentWg.Done()

}

func jobFilterGenerator(inputs *ValueInputs, possVals *AllFiltersPossibleValues, wg *sync.WaitGroup, share []JobId, workerId int) {
	for _, jobId := range share {
		for i := 0; i < rand.Intn(6)+1; i++ {
			inputs.TechSkills <- Value{
				jobId:    jobId,
				filterId: possVals.techSkills.values[rand.Intn(possVals.techSkills.len)],
			}
		}

		for i := 0; i < rand.Intn(5)+1; i++ {
			inputs.EmpTypes <- Value{
				jobId:    jobId,
				filterId: possVals.empTypes.values[rand.Intn(possVals.empTypes.len)],
			}
		}

		for i := 0; i < rand.Intn(3)+1; i++ {
			inputs.ExpLevels <- Value{
				jobId:    jobId,
				filterId: possVals.expLevels.values[rand.Intn(possVals.expLevels.len)],
			}
		}

		for i := 0; i < rand.Intn(3)+1; i++ {
			inputs.TypesOfWork <- Value{
				jobId:    jobId,
				filterId: possVals.typesOfWork.values[rand.Intn(possVals.typesOfWork.len)],
			}
		}

		for i := 0; i < rand.Intn(3)+1; i++ {
			inputs.WorkModes <- Value{
				jobId:    jobId,
				filterId: possVals.workModes.values[rand.Intn(possVals.workModes.len)],
			}
		}
	}

	wg.Done()
}

func insertFilters(tableName string, filterCol string, input []Value, wg *sync.WaitGroup) {
	wg.Add(1)
	startTime := time.Now()
	rowsCopied, err := db.DB.CopyFrom(context.Background(), pgx.Identifier{tableName}, []string{
		"job_id", filterCol,
	}, pgx.CopyFromSlice(len(input), func(i int) ([]interface{}, error) {
		return []interface{}{
			input[i].jobId, input[i].filterId,
		}, nil
	}))

	if err != nil {
		slog.Error(fmt.Sprintf("Error while COPYing into %s", tableName), "err", err)
	}

	slog.Info(fmt.Sprintf("Seeding table %s finished.", tableName),
		"time taken", time.Since(startTime),
		"rows copied", fmt.Sprintf("%d/%d", rowsCopied, len(input)))
	wg.Done()
}
