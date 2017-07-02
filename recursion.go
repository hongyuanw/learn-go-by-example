package main

import "fmt"

func fact(i int) int {

	if i == 0 {
		return 1
	}

	return i * fact(i-1)

}

func main() {

	fmt.Println(fact(3))

}
