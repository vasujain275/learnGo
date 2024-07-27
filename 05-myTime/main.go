package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to study of Time in Go")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.February, 7, 10, 10, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("02-01-2006"))
}
