package main

import "encoding/base64"
import "fmt"

func main() {

	data := []byte("demo base")
	fmt.Println(base64.StdEncoding.EncodeToString(data))

}
