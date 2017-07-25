package main

import "fmt"

// Index find string in slices
func Index(s []string, t string) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}

	return -1
}

func main() {
	var strs = []string{"peef", "bee", "ef"}

	fmt.Println(Index(strs, "bee"))
}
