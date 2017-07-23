package main

import "fmt"
import "os"

func main() {
	f := createFile("/tmp/data")

	defer closeFile(f)
	writeFile(f)

}

func createFile(filePath string) *os.File {
	fmt.Println("creating...")
	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing...")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing...")
	f.Close()
}
