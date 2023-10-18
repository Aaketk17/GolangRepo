package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myChan := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// Receiving Channel
	go func(ch <-chan int, wg *sync.WaitGroup) { //! ????????
		close(myChan) //! ????????????
		val, isChannelOpen := <-myChan

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		fmt.Println(<-myChan)
		defer wg.Done()
	}(myChan, wg)

	// Sending channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <- 5
		myChan <- 6

		// close(myChan)
		defer wg.Done()
	}(myChan, wg)

	wg.Wait()
}
