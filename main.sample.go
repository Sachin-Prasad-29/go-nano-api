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

// Model for Course - file

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

//fake Db
var courses []Course

// helper methods -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Creating a mini Api using Go Lang")
	r := mux.NewRouter()

	//seeding of fake DB
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Reactjs",
		CoursePrice: 123,
		Author: &Author{
			Fullname: "Sachin",
			Website:  "sachin.com",
		},
	})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "Vue Js",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Sachin",
			Website:  "sachin.vue",
		},
	})

	// routing
	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route
func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go Lang Nano APi</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// grap id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with given Id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send Some Data")
	}

	// what about -{}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id and convert it into string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	//append course into courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove , add with my Id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var newCourse Course
			json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId = params["id"]
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	// Send a response when id is not found
	json.NewEncoder(w).Encode("Id not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop , id, remove (index,index +1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted Successfully")
			return
			
		}
	}
	json.NewEncoder(w).Encode("No course found with Id")
	return
}
