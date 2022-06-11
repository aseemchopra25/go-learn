package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	fmt.Println(languages)
	fmt.Println("JS is short for: ", languages["JS"])
	delete(languages, "RB")
	fmt.Println(languages)

	// looping initials
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
