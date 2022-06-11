package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html?username=aces&payment=lskjlaksjd"

func main() {
	fmt.Println("URL for golangs")
	fmt.Println(myurl)
	//parsing the url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	qparams := result.Query()
	fmt.Printf("The type of qparams is: %T\n", qparams)
	fmt.Println(qparams["payment"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// reversing and making a url out of the chunks received

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "www.example.com",
		Path:    "lmao",
		RawPath: "user = aces",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
