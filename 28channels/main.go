package main

import (
	"fmt"
	"sync"
)

func main() {
	// channels allow different go routines to talk to each other instead of directly to wait group variable
	fmt.Println("All about Golang channels")

	mychan := make(chan int, 1) // buffered channel
	wg := &sync.WaitGroup{}

	// mychan <- 5
	// fmt.Println(<-mychan)

	wg.Add(2)
	//receive only channel
	go func(mychan <-chan int, wg *sync.WaitGroup) {
		// false means closed channel; true means 0 is legit
		val, isChannelOpen := <-mychan

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		fmt.Println(<-mychan)
		wg.Done()
	}(mychan, wg)
	//send only
	go func(mychan chan<- int, wg *sync.WaitGroup) {
		// mychan <- 5
		close(mychan)
		wg.Done()
	}(mychan, wg)

	wg.Wait()
}
