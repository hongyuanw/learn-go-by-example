package main

import (
	"fmt"
	"math"
	"reflect"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500

	const d = 1e10 / n
	fmt.Println(d)

	fmt.Println(reflect.TypeOf(d))
	fmt.Println(int64(d))
	fmt.Println(reflect.TypeOf(d))

	fmt.Println(reflect.TypeOf(n))
	fmt.Println(math.Sin(n))
	fmt.Println(reflect.TypeOf(n))
}
