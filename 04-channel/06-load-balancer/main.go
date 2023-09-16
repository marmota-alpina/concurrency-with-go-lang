package main

import (
	"fmt"
	"time"
)

const maxWorkers = 10_000
const maxJobs = 1_000_00

func main() {
	jobs := make(chan int)

	for i := 0; i < maxWorkers; i++ {
		go worker(i, jobs)
	}

	for i := 0; i < maxJobs; i++ {
		jobs <- i
	}
}

func worker(workerID int, jobs <-chan int) {
	for job := range jobs {
		fmt.Print("Worker ", workerID, " started job ", job, "\n")
		time.Sleep(time.Millisecond * 50)
	}
}
