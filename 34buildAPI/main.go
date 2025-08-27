package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

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
