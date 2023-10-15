package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Golang Maps section")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all langugaes is: ", languages)
	fmt.Println("Js is shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("After deleted Ruby", languages)

	fmt.Println("*****************************************")

	for key, value := range languages {
		fmt.Printf("Value for the key %v is %v \n", key, value)
	}

}
