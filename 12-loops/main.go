package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// Simple For Loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	for i := range days {
		fmt.Println(days[i])
	}

	rougueVaule := 3

	for rougueVaule < 10 {
		fmt.Println("Value is", rougueVaule)
		rougueVaule++
	}
}
