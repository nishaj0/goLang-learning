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
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(
		`{
			"courseName": "ReactJS Bootcamp",
			"Price": 299,
			"website": "Scrimba",
			"tags": ["web","development"]
		}`,
	)

	// var Courses course

	// checkValid := json.Valid(jsonDataFromWeb)

	// if checkValid {
	// 	fmt.Println("JSON was valid")
	// 	json.Unmarshal(jsonDataFromWeb, &Courses) // if we are not sure is this argument passing reference or copy its better to pass as reference
	// 	fmt.Printf("%#v\n", Courses)              // main.course{Name:"ReactJS Bootcamp", Price:299, Platform:"Scrimba", Password:"", Tag:[]string{"web", "development"}}
	// } else {
	// 	fmt.Println("JSON is not valid")
	// }

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// fmt.Printf("%#v\n", myOnlineData) // map[string]interface {}{"Price":299, "courseName":"ReactJS Bootcamp", "tags":[]interface {}{"web", "development"}, "website":"Scrimba"}

	for k, v := range myOnlineData {
		fmt.Printf("Key: %v value: %v type:%T\n", k, v, v)
	}
}

// * o/p
// Key: Price value: 299 type:float64
// Key: website value: Scrimba type:string
// Key: tags value: [web development] type:[]interface {}
// Key: courseName value: ReactJS Bootcamp type:string