package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web Verb Video")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "https://google.com"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(responseString.String())
	fmt.Println(byteCount)
}

func PerformPostJsonRequest() {
	// Change this to a server that actually processes/sends json
	const myurl = "https://www.google.com"

	// fake JSON payload
	requestBody := strings.NewReader(`
		{
			"name":"Aseem", 
			"skills":"golang",
			"time":"night"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
