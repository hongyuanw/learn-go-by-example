package main

import (
	"errors"
	"fmt"
	"strconv"
)

func positive(n int) (int, error) {
	if n < 0 {
		return -1, errors.New(strconv.Itoa(n) + " is not positive")
	}

	return n, nil
}

// todo define error type

func main() {
	nums := []int{1, -10}

	for _, n := range nums {
		if res, err := positive(n); err == nil {
			fmt.Println(res, " is positive")
		} else {
			fmt.Println(err)
		}
	}

}
