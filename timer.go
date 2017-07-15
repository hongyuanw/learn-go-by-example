package main

import "fmt"
import "time"

func main() {

	timer1 := time.NewTimer(time.Second * 2)

	timestamp := <-timer1.C
	fmt.Println("timer1 expired at ", timestamp)

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	// time.Sleep(time.Second * 2)
	result := timer2.Stop()

	if result {
		fmt.Println("timer2 stopped")
	}
}
