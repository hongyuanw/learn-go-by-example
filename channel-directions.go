package main

import "fmt"

func sendMsg(channel chan<- string, msg string) {
	fmt.Println("sending...")
	channel <- msg
	//<-channel  compile error : invalid operation: <-channel (receive from send-only type chan<- string):
}

func receiveMsg(channel <-chan string) {
	msg := <-channel
	fmt.Println("receive msg : ", msg)
}

func main() {

	oneChannel := make(chan string, 1)
	sendMsg(oneChannel, "hello message")
	receiveMsg(oneChannel)

}
