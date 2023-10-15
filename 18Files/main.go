package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the File handling in Golang")

	content := "This is the file creation content for section 18 of Golang series"

	file, err := os.Create("./myFirstFile.txt")
	checkNillErr(err)

	length, err := io.WriteString(file, content)
	checkNillErr(err)
	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./myFirstFile.txt")

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNillErr(err)

	fmt.Println("Content of the file is:", string(dataByte))
}
