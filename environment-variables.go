package main

import "fmt"
import "os"

func main() {

	os.Setenv("FOO", "ok")
	//FOO will be release when program exists

	fmt.Println("FOO : ", os.Getenv("FOO"))
	fmt.Println("BAR : ", os.Getenv("BAR"))

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

}
