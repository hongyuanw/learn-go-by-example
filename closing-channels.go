package main

import "fmt"

func main() {
	jobs := make(chan int, 10)
	done := make(chan bool)

	go func() {
		for {
			msg, more := <-jobs
			if more {
				fmt.Println("received job ", msg)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return // if don't return, for will always run until main goroutine end
			}
		}
	}()

	for i := 0; i <= 5; i++ {
		jobs <- i
		fmt.Println("sent job ", i)
	}
	close(jobs)
	<-done

}
