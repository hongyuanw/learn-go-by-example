package main

import "fmt"
import "time"

func worker(notify chan bool) {

	fmt.Println("into worker goroutine")
	time.Sleep(2 * time.Second)

	notify <- true
}

func main() {

	notifyChannel := make(chan bool, 2)

	go worker(notifyChannel)

	if ok := <-notifyChannel; ok {
		fmt.Println("worker finish")
	}

}
