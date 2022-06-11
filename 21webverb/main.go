package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Verb Video")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// fmt.Println("Status code: ", response.StatusCode)
	// fmt.Println("Content Length is: ", response.ContentLength)
	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	// var responseString strings.Builder
	// byteCount, _ := responseString.Write(content)
	// fmt.Println(responseString.String())
	// fmt.Println(byteCount)
}

func PerformPostJsonRequest() {
	// Change this to a server that actually processes/sends json
	const myurl = "http://localhost:8000/post"

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

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdate
	data := url.Values{}
	data.Add("firstname", "aseem")
	data.Add("lastname", "chopra")
	data.Add("email", "aseemch")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
