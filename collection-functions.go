package main

import "fmt"
import "strings"

// Index find string in slices
func Index(s []string, t string) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}

	return -1
}

func Include(s []string, t string) bool {
	return Index(s, t) >= 0
}

func Filter(s []string, f func(string) bool) []string {
	tmps := make([]string, 0)

	for _, str := range s {
		if f(str) {
			tmps = append(tmps, str)
		}
	}

	return tmps
}

func main() {
	var strs = []string{"peef", "bee", "ef"}

	fmt.Println(Index(strs, "bee"))
	fmt.Println(Include(strs, "de"))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "b")
	}))
}
