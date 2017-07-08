package main

import "fmt"

func main() {

	message := make(chan string, 10) //buffered channel

	message <- "buffered"
	message <- "message"

	fmt.Println(<-message)
	fmt.Println(<-message)

}
