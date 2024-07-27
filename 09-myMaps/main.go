package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RS"] = "Rust"
	languages["GO"] = "Golang"

	fmt.Println("List of all langs -", languages)
	fmt.Println("GO is short for - ", languages["GO"])

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
