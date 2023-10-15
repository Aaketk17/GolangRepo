package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang section")

	// days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	//? Type 01
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//? Type 02
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//? Type 03
	// for index, value := range days {
	// 	fmt.Printf("Index is %v and value is %v \n", index, value)
	// }

	//? Type 04
	roughValue := 1

	for roughValue < 10 {

		if roughValue == 5 {
			roughValue++
			continue
		}

		if roughValue == 4 {
			break
		}

		if roughValue == 9 {
			goto label
		}

		fmt.Println("Value is:", roughValue)
		roughValue++
	}

label:
	fmt.Println("Jumping to the label")

}

//! Output of the code is bit VAGUE
// Welcome to Loops in Golang section
// Value is: 1
// Value is: 2
// Value is: 3
// Jumping to the label
