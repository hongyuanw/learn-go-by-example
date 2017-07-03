package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Henry", age: 30})

	fmt.Println(person{name: "fred"})

	s := person{name: "ann", age: 39}
	fmt.Println(&s)

	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age) //automatically dereferences

	s.age = 10
	fmt.Println(s.age)
}
