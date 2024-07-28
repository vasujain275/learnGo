package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in Go")
	content := "This need to go to a file"

	file, err := os.Create("./myGoFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is: ", length)
	file.Close()
	readFile("./myGoFile.txt")
}

func readFile(filepath string) {
	databyte, err := os.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))
}
