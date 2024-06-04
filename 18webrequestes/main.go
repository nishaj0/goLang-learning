package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://httpbin.org/"

func main() {
	fmt.Println("welcome to web reqs")

	response, err := http.Get(url)
	checkNilError(err)
	defer response.Body.Close() // close response at last

	fmt.Printf("type of response: %T", response) // type of response: *http.Response

	dataBytes, err := io.ReadAll(response.Body)
	checkNilError(err)

	content := string(dataBytes)
	fmt.Println("data inside the website :", content) // data inside the website :<html lang="en">...</html>

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
