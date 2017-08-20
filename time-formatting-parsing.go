package main

import "fmt"
import "time"
import "github.com/metakeule/fmtdate"

func main() {
	t := time.Now()

	fmt.Println(t.Format(time.RFC3339))

	date := fmtdate.Format("YYYY-MM-DD hh:mm:ss", t)
	fmt.Println(date)

	//fmtdate is wonderful
}
