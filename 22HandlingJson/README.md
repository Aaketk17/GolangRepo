```
package main

import (
"encoding/json"
"fmt"
)

type course struct {
Name string `json:"coursename"`
Price int
Platform string `json:"website"`
Password string `json:"-"` //this alias will omit the password field in the response
Tags []string `json:"tags,omitempty"` //this will omit the tags field if tags is empty careful there shouldn't be space between `json:"tags, omitempty"`
}

//\* Output with the alias will be
// [
// {
// "coursename": "React",
// "Price": 100,
// "website": "LCO",
// "tags": [
// "web",
// "frontend"
// ]
// },
// {
// "coursename": "Node",
// "Price": 130,
// "website": "LCO",
// "tags": [
// "backend",
// "server"
// ]
// },
// {
// "coursename": "Terraform",
// "Price": 450,
// "website": "LCO"
// }
// ]

func main() {
fmt.Println("Welcom to Handling JSON in Golang section")
JsonEncoding()
}

func JsonEncoding() {
favCourses := []course{
{"React", 100, "LCO", "123abc", []string{"web", "frontend"}},
{"Node", 130, "LCO", "rtver", []string{"backend", "server"}},
{"Terraform", 450, "LCO", "34ngd", nil},
}

    //package this data as JSON data

    // * Matshal is use to implement the JSON
    finalJson, err := json.Marshal(favCourses)
    checkNillErr(err)
    //! output
    // [{"Name":"React","Price":100,"Platform":"LCO","Password":"123abc","Tags":["web","frontend"]},{"Name":"Node","Price":130,"Platform":"LCO","Password":"rtver","Tags":["backend","server"]},{"Name":"Terraform","Price":450,"Platform":"LCO","Password":"34ngd","Tags":null}]

    //* MarshalIntent will arrange the JSON values using the tab ("\t")
    finalJsonIndented, err := json.MarshalIndent(favCourses, "", "\t")
    checkNillErr(err)
    //! output with json.MarshalIndent(favCourses, "", "\t") (with "" prefix)
    //? without ALIAS
    // 	[
    //         {
    //                 "Name": "React",
    //                 "Price": 100,
    //                 "Platform": "LCO",
    //                 "Password": "123abc",
    //                 "Tags": [
    //                         "web",
    //                         "frontend"
    //                 ]
    //         },
    //         {
    //                 "Name": "Node",
    //                 "Price": 130,
    //                 "Platform": "LCO",
    //                 "Password": "rtver",
    //                 "Tags": [
    //                         "backend",
    //                         "server"
    //                 ]
    //         },
    //         {
    //                 "Name": "Terraform",
    //                 "Price": 450,
    //                 "Platform": "LCO",
    //                 "Password": "34ngd",
    //                 "Tags": null
    //         }
    // ]

    //! output with json.MarshalIndent(favCourses, "lco", "\t") (with "lco" prefix)
    //? without ALIAS
    // 	[
    // lco     {
    // lco             "Name": "React",
    // lco             "Price": 100,
    // lco             "Platform": "LCO",
    // lco             "Password": "123abc",
    // lco             "Tags": [
    // lco                     "web",
    // lco                     "frontend"
    // lco             ]
    // lco     },
    // lco     {
    // lco             "Name": "Node",
    // lco             "Price": 130,
    // lco             "Platform": "LCO",
    // lco             "Password": "rtver",
    // lco             "Tags": [
    // lco                     "backend",
    // lco                     "server"
    // lco             ]
    // lco     },
    // lco     {
    // lco             "Name": "Terraform",
    // lco             "Price": 450,
    // lco             "Platform": "LCO",
    // lco             "Password": "34ngd",
    // lco             "Tags": null
    // lco     }
    // lco]

    fmt.Printf("%s\n", finalJson)
    fmt.Printf("%s\n", finalJsonIndented)

}

func checkNillErr(err error) {
if err != nil {
panic(err)
}
}


```
