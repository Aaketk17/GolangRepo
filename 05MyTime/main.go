package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time package of Golang")

	currentTime := time.Now()
	fmt.Println("Direct value", currentTime)

	fmt.Println("Current time is: ", currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(2023, time.October, 15, 07, 56, 50, 40, time.Local)
	fmt.Println("Created Time: ", createdTime)
	fmt.Println("Formatted created time: ", createdTime.Format("01-02-2006 15:04:05 Monday"))
}
