package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to more json")

	EncodeJson()
}

func EncodeJson() {
	Courses := []course{
		{"ReactJS Bootcamp", 299, "Scrimba", "password", []string{"web", "development"}},
		{"NextJs Master class", 399, "JSM", "password", []string{"web", "development"}},
		{"MongoDB Bootcamp", 199, "Udemy", "password", nil},
	}

	// package this data as JSON data

	finalJSON, err := json.MarshalIndent(Courses, "", "\t")
	checkNilError(err)

	fmt.Print(string(finalJSON))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
