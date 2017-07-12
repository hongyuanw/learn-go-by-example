package main

import "fmt"

func main() {

	c1 := make(chan string, 1)
	c2 := make(chan bool)

	select {
	case msg := <-c1:
		fmt.Println("receive message: ", msg)
	default:
		fmt.Println("no message received")
	}

	select {
	case c1 <- "one messge": // c1 has one size buffer, if not, these case will be blocked
		fmt.Println("send message")
	default:
		fmt.Println("no message sent")
	}

	<-c1
	select {
	case msg := <-c1:
		fmt.Println("receive msg from c1", msg)
	case c2 <- true:
		fmt.Println("send message to c2")

	default:
		fmt.Println("no activty")
	}

}
