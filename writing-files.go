package main

import "fmt"
import "io/ioutil"
import "os"

func main() {

	f, err := os.OpenFile("/tmp/1.txt", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("writw")
	if err != nil {
		panic(err)
	}
	f.Close()

	content, err := ioutil.ReadFile("/tmp/1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)

}
