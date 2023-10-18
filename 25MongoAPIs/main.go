package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aaketk17/MongoAPIs/router"
)

func main() {
	fmt.Println("Welcome to MongoAPI section of Golang")
	fmt.Println("Server is getting started....")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running at port 4000")
}
