package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// no inheritance in golang; No super or parent

	vasu := User{"Vasu", "mail@vasujain.me", true, 19}
	fmt.Println(vasu)
	fmt.Printf("Vasu's Details are - %+v\n", vasu)
}

// Rememeber all first capital letters means they are public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
