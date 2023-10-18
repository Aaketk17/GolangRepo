package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signal = []string{"test"}

func main() {
	fmt.Println("Welcome to Goroutines section og golang")
	// go greeter("Helllo")
	// greeter("World")

	websites := []string{
		"https://www.google.com",
		"https://www.fb.com",
		"https://www.github.com",
		"https://go.dev/",
		"https://www.lco.dev",
	}

	for _, val := range websites {
		go getStatusCode(val)
		wg.Add(1)
	}
	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	mut.Lock()
	signal = append(signal, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}

/*
gorouties are used to handle the concorrency. If the keyword go is used infront of the function that means it will allocate a separate thread to execeute that function
but here once we created a thread we need to wait until it complete the task so that we need sleep function but its not a reliable method so that we are using
sync module's WaitGroup. We we create wg.Add(1) one thread and inside that function we are using defer wg.Done() so once its done it will end tha main method. wg.Wait()
is in the main method to avoid complete executing main function until thread complete its task

mutex we are using to lock the memory. sometime there may be multiple thread which will do read and write on the same memory location to handle that mutex will lock that
memory location until one goroutine finishes its operation. There is another Mutex sync.RWMutex ReadWrite Mutex it will allow all the goroutine to read the memory location but if there is
any goroutine want to write in that memory location it will kick all the goroutines which are doing read operations and locak the memory to leave it to write in that memory
*/
