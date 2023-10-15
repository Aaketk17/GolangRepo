package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Golang section of Arrays")

	var fruitList [4]string

	fruitList[0] = "Banana"
	fruitList[1] = "Orange"
	fruitList[3] = "Apple"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Lenght of Fruit list is: ", len(fruitList))

	var ages = [5]int{3, 4, 5, 6}

	fmt.Println("Ages are: ", ages)
	fmt.Println("Length of ages array is: ", len(ages))

}
