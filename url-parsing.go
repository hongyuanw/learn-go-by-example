package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "pastgres://user:pass@host.com:8000/path?location=shanghai#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)

	vals, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(vals)
	fmt.Println(vals["location"][0])

}
