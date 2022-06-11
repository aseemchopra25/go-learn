package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fruitList = append(fruitList, "Mango", "Banana")
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 123
	highScores[3] = 555
	// highScores[4] = 777

	// append reallocates the memory
	highScores = append(highScores, 555, 666, 444)
	fmt.Println(highScores)

	// Sorting
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from a slice based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
