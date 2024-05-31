package main

import "fmt"

func main() {
	fmt.Println("welcome to maps")

	var languages = make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println(languages)       // map[js:javascript py:python rb:ruby]
	fmt.Println(languages["js"]) // javascript

	delete(languages, "rb")
	fmt.Println(languages) // map[js:javascript py:python]

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
