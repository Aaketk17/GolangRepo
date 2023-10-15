package main

import "fmt"

// * Public variable 'cause L is caps in LoginToken
const LoginToken string = "somerandomValuesFromTheToken"

//! Can't use Walrus operator outside of the main function
//! numberOfUser := 3000.34353655757

func main() {
	fmt.Println("Variables Section")

	var username string = "Athavan"
	fmt.Println(username)
	fmt.Printf("Variable type is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable type is %T \n", smallVal)

	var smallFloat float64 = 255.45876798769878697
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is: %T \n", anotherVariable)

	//? Implicit declaration of variable, here we haven't mentioend the type lexter will indetify it
	var website = "www.athavantheivendram.com"
	fmt.Println(website)

	//? walrus operator
	numberOfUser := 3000.34353655757
	fmt.Println(numberOfUser)

	//* Public variable used within the main funtion
	fmt.Println(LoginToken)
	fmt.Printf("Variable type is: %T \n", LoginToken)
}
