package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Page   int `json:"page"`
	Fruits []string
}

func main() {

	resp := &response{
		Page:   1,
		Fruits: []string{"apple", "banana"},
	}

	str, _ := json.Marshal(resp)

	fmt.Println(string(str))
}
