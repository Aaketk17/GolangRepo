package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the If-Else statements in Golang")
	fmt.Println("Please enter any integer")

	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error in taking input from the user %v", err)
	} else {
		inputNumber, _ := strconv.ParseInt(strings.TrimSpace(inputString), 10, 64)
		if inputNumber%2 == 0 {
			fmt.Printf("Entered value %v is an EVEN number.", inputNumber)
		} else {
			fmt.Printf("Entered value %v is a ODD number.", inputNumber)
		}
	}
}
