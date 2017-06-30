package main

import "fmt"

func main() {

	nums := []int{2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "battel"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for k := range kvs {
		fmt.Println("key : ", k)
	}

	for i, c := range "golang" {
		fmt.Print(i, c, string(c), " ")
		fmt.Printf("%d : %c \n", i, c)
	}
}
