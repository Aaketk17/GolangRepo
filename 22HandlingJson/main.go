package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //this alias will omit the password field in the response
	Tags     []string `json:"tags,omitempty"` //this will omit the tags field if tags is empty careful there shouldn't be space between `json:"tags, omitempty"`
}

func main() {
	fmt.Println("Welcom to Handling JSON in Golang section")
	// JsonEncoding()
	JsonDecoding()
}

func JsonEncoding() {
	favCourses := []course{
		{"React", 100, "LCO", "123abc", []string{"web", "frontend"}},
		{"Node", 130, "LCO", "rtver", []string{"backend", "server"}},
		{"Terraform", 450, "LCO", "34ngd", nil},
	}

	//package this data as JSON data

	finalJson, err := json.Marshal(favCourses)
	checkNillErr(err)

	finalJsonIndented, err := json.MarshalIndent(favCourses, "", "\t")
	checkNillErr(err)

	fmt.Printf("%s\n", finalJson)
	fmt.Printf("%s\n", finalJsonIndented)
}

func JsonDecoding() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "React",
			"Price": 100,
			"website": "LCO",
			"tags": ["web","frontend"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was invalid")
	}

	// some cases where you just want to add data to key value
	fmt.Println("*************************************")

	//? Here to define the map we surely know that the key will be string but we aren't sure about the type of the value thats why we used interface{} to
	//? make it possible to handle any type of values

	var myOnlineData map[string]interface{}

	//? here reason for using &myOnlineData is to avoid using the copy of the real variable now we can update the exact real variable
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
