package main

import "fmt"

// walrus operators not allowed outside methods
var jwtToken int = 300000

// capital L means public variable
const LoginToken string = "jsaidjasdjasoidj"

func main() {
	var username string = "Aseem"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.123421423213213234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "chopraseem.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
