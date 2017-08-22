package main

import "fmt"
import "strconv"

func main() {
	f, _ := strconv.ParseFloat("12.33", 64)
	fmt.Println(f)

	i, _ := strconv.Atoi("12")
	fmt.Println(i)

	t := strconv.FormatInt(456, 10)
	fmt.Printf("%T, %v\n", t, t)
}
