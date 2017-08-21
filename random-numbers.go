package main

import "fmt"
import "math/rand"

func main() {

	fmt.Println(rand.Intn(10))

	rand.Seed(199)
	fmt.Println(rand.Intn(10))

	rand.Seed(199)
	fmt.Println(rand.Intn(10))

	r := rand.New(rand.NewSource(199))
	fmt.Println(r.Intn(10))

}
