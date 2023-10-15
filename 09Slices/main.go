package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the Golang Slices section")

	var fruitList = []string{"Apple", "Orange", "Banana"}
	fmt.Printf("Type of the fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 5)

	highScores[0] = 12
	highScores[1] = 1245
	highScores[2] = 45
	highScores[3] = 65
	highScores[4] = 10

	//! This will give an error as the length is defined as 4. (Already memory is allocated for 4 values)
	//! highScores[5] = 34

	fmt.Println(highScores)

	//* Here it will not throw error becuase append will reallocate the memory and adjust as per the new addition of the values
	highScores = append(highScores, 343, 73, 795)
	fmt.Println(highScores)

	fmt.Println("Before Sorting", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println("After Sorting", sort.IntsAreSorted(highScores))

	fmt.Println("***********************************************************")
	fmt.Println()
	fmt.Println()

	var courses = []string{"ReactJS", "NodeJS", "Terraform", "Golang", "Java", "Javascript", "Bash"}
	fmt.Println(courses)

	index := 4

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
