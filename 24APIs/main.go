package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// create fake db
var courses []Course

// middleware or helper files
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to writing APIs in Golang")

	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 200, Author: &Author{Fullname: "Athavan", Website: "athavan.go"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Golang", CoursePrice: 150, Author: &Author{Fullname: "Adithya", Website: "adithya.go"}})

	r.HandleFunc("/home", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/oneCourse/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/createCourse", createOneCourse).Methods("POST")
	r.HandleFunc("/updateCourse/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/deleteCourse/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

//? controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Athavan</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "applicatiopn/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatiopn/json")

	// get if from request
	params := mux.Vars(r)
	fmt.Printf("Params values %v\n", params)

	// loop through course and find the matching if and return the responses
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatiopn/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body of request is empty")
	}

	// if body is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the body")
		return
	}

	//? generate unique id and convert to string

	// Create a new source with a specific seed value
	source := rand.NewSource(time.Now().UnixNano())

	// Create a new random number generator
	rng := rand.New(source)

	course.CourseId = strconv.Itoa(rng.Intn(100))

	//? append course into courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// get id from the req
	params := mux.Vars(r)

	// loop through value and find the course with the id and remove the values and add new values with the same id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO: Should handle of the id is empty

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Loop and get the course to delete
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is deleted")
			break
		}
	}

}
