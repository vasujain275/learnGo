package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Study of Slices in GO")

	var fruitList = []string{"Apple", "Peach"}
	fmt.Printf("Type of Fruit list is: %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	// Slicing the slice
	fruitList = append(fruitList[1:3]) // Ranges are inclusive i.e. 1,3 will not be added
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 232
	highScores[1] = 843
	highScores[2] = 941
	highScores[3] = 365

	highScores = append(highScores, 854, 861, 992)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// How to Remove a Value from Slice Based on Index

	var courses = []string{"svelte", "javascript", "java", "python", "rust"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
