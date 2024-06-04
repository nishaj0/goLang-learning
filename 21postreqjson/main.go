package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to get request")

	const myUrl string = "http://localhost:8000/post"

	// PerformPostJsonRequest(myUrl)
	PerformPostFormRequest(myUrl)
}

func PerformPostJsonRequest(reqUrl string) {

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename":"golang",
		"price":0,
		"platform":"website"
	}
	`)

	response, err := http.Post(reqUrl, "application/json", requestBody)
	checkNilError(err)

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content)) // {"coursename":"golang","price":0,"platform":"website"}
}

func PerformPostFormRequest(reqUrl string) {
	data := url.Values{}

	data.Add("firstname", "John")
	data.Add("lastname", "Doe")
	data.Add("email", "john@go.dev")

	response, err := http.PostForm(reqUrl, data)
	checkNilError(err)
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content)) // {"email":"john@go.dev","firstname":"John","lastname":"Doe"}
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

