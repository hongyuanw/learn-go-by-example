package main

import "flag"
import "fmt"
import "os"

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("numb", 11, "an int")

	var boo bool
	flag.BoolVar(&boo, "fork", false, "a bool")

	flag.Parse()

	fmt.Println("word : ", *wordPtr)
	fmt.Println("num : ", *numPtr)
	fmt.Println("bool : ", boo)
	fmt.Println("bool var : ", boo)

	fmt.Println("tail : ", flag.Args())

	fmt.Println("os args : ", os.Args)

}
