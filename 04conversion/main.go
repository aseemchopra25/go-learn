package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our Pizza between 1-5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", strings.TrimSpace(input))
	numRating, err := strconv.Atoi(strings.TrimSpace(input))
	fmt.Println(numRating)
	fmt.Println("Added 1 to your rating: ", numRating+1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No errors encountered!")
	}
}
