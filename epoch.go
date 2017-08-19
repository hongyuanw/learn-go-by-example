package main

import "fmt"
import "time"

func main() {

	now := time.Now()

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	fmt.Println(now)
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))

}
