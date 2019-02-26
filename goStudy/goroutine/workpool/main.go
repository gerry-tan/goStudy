package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id     int
	Number int
}

type Result struct {
	job *Job
	sum int
}

func calc(job *Job, result chan *Result) {
	number := job.Number

	var sum int
	for number != 0 {
		temp := number % 10
		sum += temp
		number /= 10
	}

	r := &Result{
		job: job,
		sum: sum,
	}

	result <- r
}

func startWorkPool(num int, jobs chan *Job, results chan *Result) {
	for i := 0; i < num; i++ {
		go worker(jobs, results)
	}
}

func worker(jobs chan *Job, result chan *Result) {
	for job := range jobs {
		calc(job, result)
	}
}

func printResult(results chan *Result) {
	for rs := range results {
		fmt.Printf("result jobId=%d, number=%d, sum=%d\n", rs.job.Id, rs.job.Number, rs.sum)
	}
}

func main() {
	jobChan := make(chan *Job, 1000)
	resultChan := make(chan *Result, 1000)

	startWorkPool(10, jobChan, resultChan)

	go printResult(resultChan)

	for id := 1; id <= 1000; id++ {
		number := rand.Int()
		job := &Job{
			Id:     id,
			Number: number,
		}
		jobChan <- job
	}
}
