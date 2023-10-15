package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers Section")

	var ptr *int
	fmt.Println("Value of the pointer is: ", ptr)

	randomNumber := 453

	//* creating a pointer which is referening to some memory location
	var rndPtr = &randomNumber

	fmt.Println("Memory address of the randomNumber variable is: ", rndPtr)
	fmt.Println("Value of the randomNumber is: ", *rndPtr)

	*rndPtr = *rndPtr + 7
	fmt.Println("New value of the randomNumber variable is: ", randomNumber)
}
