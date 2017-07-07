package main

import "fmt"

func printStr(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {

	printStr("golang")

	go printStr("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
