package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to handling web requests in Golang")

	response, err := http.Get(url)
	checkNillErr(err)
	fmt.Printf("Type of response is %T\n", response) //? Type of response is *http.Response, So its a pointer and its useful in modifications
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	checkNillErr(err)

	content := string(dataBytes)
	fmt.Println(content)
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
