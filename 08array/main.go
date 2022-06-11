package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	// Have to provide a number in array
	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit list is", fruitlist)
	fmt.Println("Fruit list is", len(fruitlist))

	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("VegeList is: ", vegList)
}
