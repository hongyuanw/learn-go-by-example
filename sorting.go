package main

import "fmt"
import "sort"

func main() {
	strs := []string{"a", "e", "c"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{1, 3, 2}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted : ", s)
}
