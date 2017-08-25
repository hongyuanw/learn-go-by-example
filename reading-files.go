package main

import "fmt"
import "io/ioutil"

func main() {
	content, err := ioutil.ReadFile("/tmp/1.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)

}
