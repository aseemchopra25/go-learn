package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File operations in Go")
	content := "This needs to go in a file -- Golang"
	file, err := os.Create("./myfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./myfile.txt")

}

func readFile(filename string) {
	// Not being read in the string format
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	// fmt.Println("Text data inside the file is: ", databyte)
	fmt.Println("Text data inside the file is: \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
