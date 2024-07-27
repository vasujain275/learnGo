package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in Go")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("FruitList is :", fruitList)
	fmt.Println("FruitList length is :", len(fruitList)) // Will always return 4 i.e size
}
