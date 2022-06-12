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

type Course struct {
	CourseID    string  `json:"courseid"`
	Coursename  string  `json:"coursename"`
	CoursePrice int     `json:"-"` // hides price if you put hyphen else put price
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db
var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseID == "" && c.Coursename == ""
	return c.Coursename == ""

}

func main() {

	fmt.Println("--------Welcome to a Golang API------------")
	r := mux.NewRouter()

	// seeding

	courses = append(courses, Course{CourseID: "2", Coursename: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Aseem Chopra", Website: "chopraaseem.com"}})
	courses = append(courses, Course{CourseID: "4", Coursename: "Golang", CoursePrice: 199, Author: &Author{Fullname: "Aseem Chopra", Website: "chopraaseem.com"}})

	//routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// Controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)
	fmt.Println(params)

	// loop through the courses; find matching idea and return the response

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return

		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	for _, c := range courses {
		if c.Coursename == course.Coursename {
			json.NewEncoder(w).Encode("Error! Course with the same name already exists!")
			return
		}
	}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id, string
	// append this new course into courses

	rand.Seed(time.Now().Unix())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	// loop through values and when we hit ID, we need to remove, add operation again but this time, add with my ID

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}

	}

	// TODO send a response when ID is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, hit id, remove operation (index, index + 1)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// Send a json response that we deleted the course
			break
		}
	}
}
