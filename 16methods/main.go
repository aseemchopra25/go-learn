package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	// no inheritance in golang lol? no super lol no classes etc.
	aseem := User{"Aseem", "aseemchopra@protonmail.com", true, 5, 12}
	fmt.Println(aseem)
	// %+v for more details, makes things easier
	fmt.Printf("Details of aseem are: %+v]\n", aseem)
	fmt.Printf("Name is %v and email is %v\n", aseem.Name, aseem.Email)
	aseem.GetStatus()
	aseem.NewMail()
	fmt.Printf("Name is %v and email is %v\n", aseem.Name, aseem.Email)
	// The difference here is due to copies being passed instead of pointers
}

// Public so Uppercase first letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)

}
