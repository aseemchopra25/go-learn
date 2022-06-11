package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	// no inheritance in golang lol? no super lol no classes etc.
	aseem := User{"Aseem", "aseemchopra@protonmail.com", true, 5}
	fmt.Println(aseem)
	// %+v for more details, makes things easier
	fmt.Printf("Details of aseem are: %+v]\n", aseem)
	fmt.Printf("Name is %v and email is %v", aseem.Name, aseem.Email)
}

// Public so Uppercase first letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
