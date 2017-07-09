package main

import "fmt"
import "time"

func main() {

	message := make(chan string) // channel

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("into another goruoutie")
		message <- "ping"
	}()

	fmt.Println("waiting ...")
	msg := <-message
	fmt.Println(msg)

}
