package main

import "fmt"
import "regexp"
import "time"

func main() {

	t := time.Now()
	for i := 0; i < 100; i++ {
		res, _ := regexp.MatchString("a([a-z]+)o", "afdjo")
		fmt.Print(res)
	}
	fmt.Println("cost without regexp compiling", time.Now().Sub(t).Nanoseconds())

	t = time.Now()
	reg := regexp.MustCompile("a([a-z]+)o")
	for i := 0; i < 100; i++ {
		rest := reg.MatchString("afdjo")
		fmt.Print(rest)
	}

	fmt.Println("cost with regexp compiling", time.Now().Sub(t).Nanoseconds())

	//cost without compiling is 10 times than that with compiling
}
