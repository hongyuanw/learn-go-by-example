package main

import "fmt"
import "time"

func main() {

	now := time.Now()

	fmt.Println(now.Weekday())

	duration := now.Sub(time.Now())

	fmt.Println(duration)
	fmt.Println(duration.Nanoseconds())

}
