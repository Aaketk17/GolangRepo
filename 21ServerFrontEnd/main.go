package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Server for golang frontend")
	// PerformGetRequest()
	// PerformPostRequest()
	// PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	checkNillErr(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Lenght ", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	checkNillErr(err)

	//* Two ways to handle the content from the above line one is using string to convert it
	fmt.Println(string(content))

	fmt.Println("****************************************************")

	//* Second is using strings package's builder
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println("Content is: ", responseString.String())
}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json pyload
	requestBody := strings.NewReader(`
		{
			"name": "Athavan",
			"Address" : "Jaffna",
			"Age": 21
		}
	`)

	reponse, err := http.Post(myUrl, "application/json", requestBody)
	checkNillErr(err)

	defer reponse.Body.Close()

	content, err := io.ReadAll(reponse.Body)
	checkNillErr(err)

	//* Method 1
	fmt.Println(string(content))

	//* Method 2
	fmt.Println("****************************************************")

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println("Content is: ", responseString.String())

}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//* create formdata
	data := url.Values{}
	data.Add("FirstName", "Athavan")
	data.Add("LastName", "Theivendram")
	data.Add("Email", "athavan@go.dev")

	response, err := http.PostForm(myUrl, data)
	checkNillErr(err)

	content, err := io.ReadAll(response.Body)
	checkNillErr(err)

	fmt.Println(string(content))
	fmt.Println("****************************************************")

	defer response.Body.Close()

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println("Content is: ", responseString.String())

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
