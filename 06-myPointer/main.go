package main

import "fmt"

func main() {
	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual Pointer is ", ptr)
	fmt.Println("Value of inside the Pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of after multiply ", *ptr)
}
