package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// no inheritance in golang; No super or parent

	vasu := User{"Vasu", "mail@vasujain.me", true, 19}
	fmt.Println(vasu)
	fmt.Printf("Vasu's Details are - %+v\n", vasu)
	vasu.GetStatus()
	vasu.NewMail()
	fmt.Printf("Vasu's Details are - %+v\n", vasu)
}

// Rememeber all first capital letters means they are public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is: ", u.Email)
}
