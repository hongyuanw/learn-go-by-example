package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		uppers := strings.ToUpper(scanner.Text())

		fmt.Println(uppers)
	}

	if err := scanner.Err(); err != nil {

		fmt.Println(err)
	}
}
