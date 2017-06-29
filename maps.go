package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 3

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, exist := m["k2"]
	fmt.Println("k2 exist :", exist)

	n := map[string]int{"foo1": 1, "foo2": 2}
	fmt.Println("n : ", n)

	m["k3"] = 0
	v3, k3Exists := m["k3"]
	v4, k4Exists := m["k4"] // not exist

	fmt.Println("v3 : ", v3, k3Exists)
	fmt.Println("v4 : ", v4, k4Exists)
}
