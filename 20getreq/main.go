package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to get request")

	const myUrl string = "http://localhost:8000/get"

	PerformGetRequest(myUrl)
}

func PerformGetRequest(reqUrl string) {
	response, err := http.Get(reqUrl)
	checkNilError(err)

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("byte count is:", byteCount)
	fmt.Println(responseString.String())
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// * o/p

// welcome to get request
// Status Code:  200  
// Content Length:  43
// byte count is: 43  
// {"message":"Hello from learnCodeonline.in"}
