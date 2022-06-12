package main

import (
	"fmt"
	"net/http"
	"sync"
)

// time.Sleep() on steroids
// they are usually pointers

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex // pointer

func main() {
	websitelist := []string{
		"https://google.com",
		"https://go.dev",
		"https://chopraaseem.com",
		"https://github.com",
		"https://fb.com",
		"https://instagram.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	//doesn't allows main to finish till everything is done
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS")
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

}
