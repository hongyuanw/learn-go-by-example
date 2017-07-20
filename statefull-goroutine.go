package main

import (
	"fmt"
	"time"
)

func main() {

	size := 100
	signal := make(chan bool, size)

	queue := make(chan int)
	go func() {
		var state int = 0

		for {
			inc, more := <-queue
			if more {
				state += inc // increment in one goroutine
			} else {
				break
			}
		}

		fmt.Println("state : ", state)

	}()

	for i := 0; i < size; i++ {
		go func() {
			for j := 0; j < 300; j++ {
				queue <- 1 // send each write operation to queue, just serialization
			}
			signal <- true
		}()

	}

	for i := 0; i < size; i++ {
		<-signal
	}

	close(queue)

	time.Sleep(time.Second)
}
