package main

import "fmt"
import "time"

func main() {

	ticker1 := time.NewTicker(time.Millisecond * 200)

	go func() {
		// regular intervals
		for timestamp := range ticker1.C {
			fmt.Println("tick time : ", timestamp)
		}
	}()

	time.Sleep(time.Second) // blocking main goroutine
	ticker1.Stop()

	fmt.Println("ticker stop")

}
