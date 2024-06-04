package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=fjdas242aeio"

func main() {
	fmt.Println("welcome to url handling")
	fmt.Println(myUrl)

	// parsing
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // loc.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.Port())   // 3000
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=fjdas242aeio

	qparams := result.Query()
	fmt.Printf("type of query params: %T\n", qparams) // type of query params: url.Values => basically they are maps in go
	fmt.Println(qparams["coursename"])                // [reactjs]

	// get all query params value
	for _, val := range qparams {
		fmt.Println("params is", val)
	}
	// params is [reactjs]
	// params is [fjdas242aeio]

	// creating url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "github.dev",
		Path:    "explore/p1",
		RawQuery: "feed=latest",
	}

	anotherurl := partsOfUrl.String()
	fmt.Println(anotherurl) // https://github.dev/explore/p1?feed=latest
}
