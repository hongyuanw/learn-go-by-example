package main

import "os"

func main() {

	// the panic will stop the main goroutine
	// and print message
	panic("a problem")

	_, err := os.Create("/tmp/one")
	if err != nil {
		panic(err)
	}
}
