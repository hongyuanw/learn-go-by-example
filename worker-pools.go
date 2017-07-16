package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	// waiting for job
	for job := range jobs {
		fmt.Println("worker ", id, "start job ", job)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finish job", job)
		results <- job
	}
}

func main() {

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	count := 7

	//init workers
	for i := 0; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	//send jobs to channel
	for i := 0; i <= count; i++ {
		jobs <- i
	}

	//	close(jobs) if jobs channel is not closed, other goroutines will be dead after main() goroutine

	for i := 0; i <= count; i++ {
		<-results
	}
}
