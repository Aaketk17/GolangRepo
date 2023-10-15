package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang's Structs sections")

	athavan := User{"Athavan", 26, "athavan@dev.go", true, 23.12}
	fmt.Println(athavan)
	fmt.Printf("Athavan's details are: %+v\n", athavan) //* Output will be "Athavan's details are: {Name:Athavan Age:26 Email:athavan@dev.go Status:true Height:23.12}"
	fmt.Printf("Name is %v and age is %v", athavan.Name, athavan.Age)

}

// ! First letter should be Capital
type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
	Height float32
}
