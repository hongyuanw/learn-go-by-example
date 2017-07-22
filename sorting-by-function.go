package main

import "fmt"
import "sort"

type ByLength []string

// https://golang.org/pkg/sort/#Interface
func (s ByLength) Len() int {
	return len(s)
}

// i, j index
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] // go swap !
}

func main() {
	strs := []string{"dem", "freestyle", "jay"}

	sort.Sort(ByLength(strs)) // type cast

	fmt.Println(strs)
}
