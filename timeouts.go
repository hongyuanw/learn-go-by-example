package main

import "fmt"
import "time"

func main() {

	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "result 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println("get ", msg)

	case <-time.After(time.Second * 2):
		// https://golang.org/pkg/time/#After  After return a channel : <-chan Time
		fmt.Println("timeout..")
	}

}
