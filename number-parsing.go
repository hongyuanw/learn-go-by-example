package main

import "fmt"
import "strconv"

func main() {
	f, _ := strconv.ParseFloat("12.33", 64)
	fmt.Println(f)

	i, _ := strconv.Atoi("12")
	fmt.Println(i)
}
