package main

import "fmt"

func main() {

	message := make(chan string) // channel

	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)

}
