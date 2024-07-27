package main

import "fmt"

func main() {
	fmt.Println("If-Else in Go")

	loginCount := 23
	result := ""

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Super User"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	}
}
