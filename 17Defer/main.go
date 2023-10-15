package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Defer section of Golang")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// ? When Defer keyword used that line will stop executing and the execution start from the next line once after all the lines are executed at last the defer keyword line will execute
// ? we can think this as a defer line will execute only at last onece after all the lines are executed

//? Where there are multiple defer once all other lines are executed the defer lines executed in the reverse order line which defer first will execute at last
