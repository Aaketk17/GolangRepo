package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Race conditions in Golang")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	// mutR := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3) //or
	//wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("First Function")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//wg.Add(2)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Second Function")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Third Function")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
	// 	fmt.Println("Third Function")
	// 	mutR.RLock()
	// 	fmt.Println(score)
	// 	mutR.RUnlock()
	// 	wg.Done()
	// }(wg, mutR)

	wg.Wait()
	fmt.Println(score)
}

/*
Best place to add the lock is at the place where we are reading or writing the resource
*/
