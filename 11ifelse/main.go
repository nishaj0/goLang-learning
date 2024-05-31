package main

import "fmt"

func main() {
	fmt.Println("welcome to if and else")

	// loginCount := 11
	var result string

	// if loginCount < 10 {
	// 	result = "regular users"
	// } else if loginCount > 10 {
	// 	result = "watch out"
	// } else {
	// 	result = "10 users"
	// }

	// if 9%2 == 0 {
	// 	result = "number is even"
	// } else {
	// 	result = "number us odd"
	// }

	if num, limit, _ := 6, 10, "number is"; num < limit {
		result = " less than 10"
	} else if num > limit {
		result = " grater than 10"
	}

	fmt.Println(result) //  less than 10
}
