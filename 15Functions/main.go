package main

import "fmt"

//! In Go we can't define function within a function

func main() {
	fmt.Println("Welcome to the Functions section of Golang")

	greeter()
	addedValue := adder(5, 6)

	fmt.Println("Results from the adder function:", addedValue)

	proAdderValue, proAdderMessage := proAdder(4, 5, 4, 90)
	fmt.Println("Results from the proAdder function:", proAdderValue)
	fmt.Println("Message from the proAdder function:", proAdderMessage)

}

func greeter() {
	fmt.Println("This is a greeter function")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "This is my message"
}
