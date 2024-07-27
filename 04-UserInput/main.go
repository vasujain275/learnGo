package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input Program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for our Pizza:")

	// comma ok || err err syntax

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Thanks for rating,", input)
}
