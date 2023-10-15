package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Golang Switch case section")

	// Create a new source with a specific seed value
	source := rand.NewSource(time.Now().UnixNano())

	// Create a new random number generator
	rng := rand.New(source)

	diceNumber := rng.Intn(6) + 1

	//! rand.Seed(time.Now().UnixNano()) This code to generate random number depreciated

	fmt.Println("Random generated value is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can now move 2 spots")
	case 3:
		fmt.Println("You can now move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can now move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can now move 5 spots")
	case 6:
		fmt.Println("You can now move 6 spots")
	default:
		fmt.Println("What's going on here!!!!")
	}
}

//* If Number 3 is generated this will be the output:
// Welcome to Golang Switch case section
// Random generated value is:  3
// You can now move 3 spots
// You can now move 4 spots
// You can now move 5 spots

//* Reason is when using keyword fallthrough it will go to the next case as well. As here we are having fallthrough in case 3 and 4 when we get number 3 it print 3,4 and 5
