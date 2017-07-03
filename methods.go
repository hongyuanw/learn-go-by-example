package main

import "fmt"

type person struct {
	age, grade int
}

func (p person) updateAge() int {
	p.age += 1

	return p.age
}

func (ptr *person) updateGrade() int {
	ptr.grade += 1

	return ptr.grade
}

func main() {

	p := person{age: 10, grade: 100}

	fmt.Println(p)

	newAge := p.updateAge()
	fmt.Println(p)
	fmt.Println("new age : ", newAge)

	ptr := &p
	newGrade := ptr.updateGrade() //value and pointer method call both use dot
	fmt.Println(ptr)
	fmt.Println("new grade : ", newGrade)
}
