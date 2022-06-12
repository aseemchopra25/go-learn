package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to math and crypto in Golang")

	// var numOne int = 2
	// var numTwo float64 = 4

	// fmt.Println("The sum is: ", numOne+int(numTwo))

	// random numbers - math (inferior algorithm), crypto (superior algorithm)

	// math based rng

	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(5))

	// crypto based random

	crypt, _ := rand.Int(rand.Reader, big.NewInt(100000))
	fmt.Println(crypt)

}
