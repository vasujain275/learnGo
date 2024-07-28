package main

import "fmt"

// defer keywords make the rest of go at the very end of function i.e. just above the closing }
// It follow the lifo model meaning last in fist out

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Output -
// Hello
// 4
// 3
// 2
// 1
// 0
// Two
// One
// World
