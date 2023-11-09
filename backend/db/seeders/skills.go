package seeders

import (
	"context"
	"fmt"
	"log"
	"main/db"

	// "log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
)

type jobId int

func SeedSkills() {
	fmt.Println("\nStart seeding job skills...")
	rows, err := db.DB.Query(context.Background(), "SELECT id FROM jobs")

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

	skills := *generateSkills(jobIds)
	start := time.Now()

	rowsCopied, err := db.DB.CopyFrom(context.Background(), pgx.Identifier{"job_skills"},
		[]string{"name", "job_id"}, pgx.CopyFromSlice(len(skills), func(i int) ([]any, error) {
			return []any{&skills[i].name, &skills[i].jobId}, nil
		}))

	if err != nil {
		log.Fatalln("Error while inserting rows", err)
	}

	fmt.Printf("Inserted %d/%d skills in %s.\n", rowsCopied, len(skills), time.Since(start))
}

type Skill struct {
	name  string
	jobId jobId
}

func generateSkills(jobIds []jobId) *[]Skill {
	var wg sync.WaitGroup

	jobIdsCount := len(jobIds)
	var jobIdsShares [][]jobId
	var workerJobShare int
	var workerCount int

	if jobIdsCount >= 1000 {
		workerJobShare = 500
	} else {
		workerJobShare = jobIdsCount / 2
	}

	indexStart := 0
	indexEnd := workerJobShare
	for i := 0; i < jobIdsCount/workerJobShare; i++ {
		jobIdsShares = append(jobIdsShares, jobIds[indexStart:indexEnd])
		indexStart = indexEnd
		indexEnd += workerJobShare
	}

	if jobIdsCount%workerJobShare != 0 {
		if indexEnd > jobIdsCount {
			jobIdsShares = append(jobIdsShares, jobIds[indexStart:])
		} else {
			jobIdsShares = append(jobIdsShares, jobIds[indexEnd:])
		}
	}

	workerCount = len(jobIdsShares)
	var skills []Skill
	const MAX_NUMBER_OF_SKILLS = 6
	totalTimeStart := time.Now()

	skillsChan := make(chan Skill, 2*workerJobShare)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go func(workerId int, idsShare []jobId, c chan<- Skill) {
			defer wg.Done()

			start := time.Now()

			for _, jobId := range idsShare {
				numbOfSkills := rand.Intn(MAX_NUMBER_OF_SKILLS) + 1

				for j := 0; j < numbOfSkills; j++ {

					// skills = append(skills, Skill{
					// 	name:  "Skill" + strconv.Itoa(j+1),
					// 	jobId: jobId,
					// })
					c <- Skill{
						name:  "Skill" + strconv.Itoa(j+1),
						jobId: jobId,
					}
				}
			}
			fmt.Printf("Worker %d finished generating skills for %d jobs in %s.\n", workerId, len(idsShare), time.Since(start))
		}(i+1, jobIdsShares[i], skillsChan)
	}
	// var wg2 sync.WaitGroup
	// wg2.Add(1)
	// go func(c <-chan Skill, count int) {
	// 	isok := true
	// 	for {
	// 		skill, ok := <-c
	// 		if !ok {
	// 			isok = false
	// 			break
	// 		}
	// 		skills = append(skills, skill)

	// 		if !isok {
	// 			wg2.Done()
	// 			break
	// 		}
	// 	}

	// 	// for skill := range c {
	// 	// 	skills = append(skills, skill)
	// 	// }

	// }(skillsChan, jobIdsCount)
	// wg2.Wait()
	go func(c chan Skill) {
		for skill := range c {
			skills = append(skills, skill)
		}
	}(skillsChan)

	wg.Wait()
	close(skillsChan)

	fmt.Printf("%d worker(s) finished generating %d skills for %d jobs in %s\n",
		workerCount, len(skills), jobIdsCount, time.Since(totalTimeStart))

	return &skills
}
