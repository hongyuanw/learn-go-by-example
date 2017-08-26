package main

import "fmt"
import "os"

func main() {
	argsWithProg := os.Args

	argsWithoutProg := os.Args[1:]

	arg := os.Args[2]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}
