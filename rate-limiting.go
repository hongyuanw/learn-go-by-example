package main

import "fmt"
import "time"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(time.Millisecond * 200)

	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}

	//allow bursts of up to 3 events
	fmt.Println("burst request ...")

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	for i := 0; i < 10; i++ {
		<-burstyLimiter
		fmt.Println("request ", i, time.Now())
	}

}
