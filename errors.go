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

type argError struct {
	arg         int
	description string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.description)
}

func positive2(n int) (int, error) {
	if n < 0 {
		return -1, &argError{arg: n, description: "positive2 error"}
	}

	return n, nil
}

func main() {
	nums := []int{1, -10}

	for _, n := range nums {
		if res, err := positive(n); err == nil {
			fmt.Println(res, " is positive")
		} else {
			fmt.Println(err)
		}
	}

	_, e := positive2(-1)
	if ae, ok := e.(*argError); ok { // type assertion
		fmt.Println(ae.arg)
		fmt.Println(ae.description)
		fmt.Println(ae.Error())
	}
}
