package main

import "fmt"
import "os"

func main() {
	defer fmt.Println("!") // will never run

	os.Exit(1)
}
