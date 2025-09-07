package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	// mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("1 go routing")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("2 go routing")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("3 go routing")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("3 go routing")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}

// go run --race .
