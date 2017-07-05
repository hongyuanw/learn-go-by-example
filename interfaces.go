package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	c := circle{radius: 1.9}

	fmt.Println(c.area())
	measure(&c)

}
