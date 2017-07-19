package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var num1 uint64 = 0
	var num2 uint64 = 0
	var num3 uint64 = 0

	size := 100
	signal := make(chan bool, size)

	var mutex = &sync.Mutex{}

	for i := 0; i < size; i++ {
		go func() {
			for j := 0; j < 300; j++ {
				num1 += 1

				mutex.Lock()
				num2 += 1
				mutex.Unlock()

				atomic.AddUint64(&num3, 1)
			}
			signal <- true
		}()

	}

	for i := 0; i < size; i++ {
		<-signal
	}

	fmt.Println("num1 ", num1)
	fmt.Println("num2 ", num2)
	fmt.Println("num3 ", num3)

	//num1  28642   invalid num caused by data race
	//num2  30000
	//num3  30000
}
