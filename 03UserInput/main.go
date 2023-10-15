package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the Pizza")

	//? comma ok || err ok format
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving rating, ", input)
	fmt.Printf("Type of the user input is %T", input)
}
