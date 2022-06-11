package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter()
	result := adder(2, 3)
	result, lol := proAdder(1, 2, 3, 4, 5, 6, 6, 7)
	fmt.Println(result, lol)
}

func adder(valOne int, valTwo int) int { //this is a function signature at the end
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "lmao"
}

func greeter() {
	fmt.Println("Hey man!")
}

// Anonymous functions might not have a name and be called immediately
