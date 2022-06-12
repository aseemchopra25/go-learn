package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions in Golang")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}
	// rlock/runlock over println in go routine;
	//strict lock not necessary for reading; we can allow reading threads
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		mut.Lock()
		fmt.Println("One Routine")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		mut.Lock()
		fmt.Println("Two Routine")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		mut.Lock()
		fmt.Println("Three Routine")
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}
