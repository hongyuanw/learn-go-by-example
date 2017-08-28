package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)

	done := make(chan bool)

	signal.Notify(sig, syscall.SIGINT)

	go func() {
		oneSig := <-sig
		fmt.Println()
		fmt.Println(oneSig)
		done <- true

	}()

	fmt.Println("awaiting signal...")
	<-done

	fmt.Println("done")

}
