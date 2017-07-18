package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops uint64 = 0

	for i := 0; i <= 10; i++ {
		go func() {

			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond * 100) // if no sleeping, atomic.LoadUint64 will almost never get value
			}
		}()
	}

	time.Sleep(time.Second * 1)

	opsValue := atomic.LoadUint64(&ops)
	fmt.Println(opsValue)

}
