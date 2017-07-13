package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // queue must be closed, otherwise range will never end -> deadlock

	for msg := range queue {
		fmt.Println(msg)
	}

}
