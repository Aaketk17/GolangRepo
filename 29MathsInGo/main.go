package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Channels in Golang")

	// var myNumberOne int = 2
	// var myNumberTwo float32 = 2

	// fmt.Println("The sum is:", myNumberOne+int(myNumberTwo))

	// Random numbers
	// Create a new source with a specific seed value
	// source := rand.NewSource(time.Now().UnixNano())

	// Create a new random number generator
	// rng := rand.New(source)
	// fmt.Println(rng.Intn(6))

	// random from crypto

	rnNo, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(rnNo)

}
