package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:300/learn?coursename=golang&paymentid=riioj46n"

func main() {
	fmt.Println("Welcome to URLs section in Golang")

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of the query params are: %T\n ", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutes",
		RawPath: "coursename=nodejs",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("New constructed URL is: ", anotherUrl)

}
