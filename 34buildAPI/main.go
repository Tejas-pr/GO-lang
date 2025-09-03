package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// helper / middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	// serveHome()
}

// controllers - file
// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the HOME</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(courses)
}

// To get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// grab the id from request
	params := mux.Vars(r)

	// loop through courses and find the id that passed
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course id !!!")
	return
}

// to create a course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one")
	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please some data !!")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&courses)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data !!")
		return
	}
	// generate UUID, string to int
	// Create a new random source with a specific seed value
	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// update
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	// first grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		} else {
			json.NewEncoder(w).Encode("course id not found !!!")
		}
	}
}
