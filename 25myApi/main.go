package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 249, Author: &Author{FullName: "John Doe", Website: "udemy.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "NodeJs", CoursePrice: 199, Author: &Author{FullName: "Jane Doe", Website: "udemy.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Golang", CoursePrice: 299, Author: &Author{FullName: "John Doe", Website: "udemy.com"}})
	courses = append(courses, Course{CourseId: "5", CourseName: "Python", CoursePrice: 199, Author: &Author{FullName: "Jane Doe", Website: "udemy.com"}})
	courses = append(courses, Course{CourseId: "6", CourseName: "Django", CoursePrice: 299, Author: &Author{FullName: "John Doe", Website: "udemy.com"}})

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(handlePageNotFound)

	fmt.Println("server started")

	log.Fatal(http.ListenAndServe(":3000", r))
}

// controller functions

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to course api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// take course id from params
	params := mux.Vars(r)

	// loop through courses, find course by id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("id not match")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
		return
	}

	// if json is empty
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("json is empty")
		return
	}

	// if course title is same inside courses
	for _, courseElement := range courses {
		if course.CourseName == courseElement.CourseName {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode("duplicate course")
			return
		}
	}

	// generate unique id as string
	// append course in to courses
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab the id from req
	params := mux.Vars(r)

	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("id not found")
		return
	}

	// loop through courses, remove the course, add the course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// remove the course
			courses = append(courses[:index], courses[index+1:]...)

			// get user sended course
			var reqCourse Course
			json.NewDecoder(r.Body).Decode(&reqCourse)

			// add id to course
			reqCourse.CourseId = params["id"]

			// append updated course to courses
			courses = append(courses, reqCourse)

			// send response
			json.NewEncoder(w).Encode(course)
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("id not found")
		return
	}

	for index, course := range courses {
		if params["id"] == course.CourseId {
			// remove course from courses
			courses = append(courses[:index], courses[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func handlePageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("page not found")
}
