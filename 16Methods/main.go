package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods section of Golang")

	athavan := User{"Athavan", 26, "athavan@dev.go", true, 23.12}

	athavan.GetStatus()
	athavan.ChangeMain()
	//* Email won't be changed in the original variable becuase change heppned only in the copy of the variable if we want to change that then we want to use the pointers
	fmt.Printf("Name is %v and email is %v \n", athavan.Name, athavan.Email)

}

// ! First letter should be Capital
type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
	Height float32
	//! If we use small letter that property won't be exported
	//! address string
}

// * Method name should start with captial letter otherwise method won't be exported
func (u User) GetStatus() {
	fmt.Println("User status is: ", u.Status)
}

func (u User) ChangeMain() {
	u.Email = "test@go.dev"
	fmt.Println("Email has been changed to ", u.Email)
}
