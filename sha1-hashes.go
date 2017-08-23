package main

import "fmt"
import "crypto/sha1"

func main() {

	s := []byte("sha1")
	fmt.Printf("%x", sha1.Sum(s))
}
