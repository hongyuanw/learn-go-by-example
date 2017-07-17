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

}
