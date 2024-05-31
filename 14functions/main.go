package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")
	// greeter()

	// result := adder(3, 1)
	result, adderString, adderString2 := adder(2, 3)
	fmt.Println(result)
	fmt.Println(adderString, adderString2)
}

// func greeter() {
// 	fmt.Println("hello from GoLang")
// }

func adder(valOne int, valTwo int) (int, string, string) {
	return valOne + valTwo, "its adder", "its adder once again"
}

// func proAdder(values ...int) int {
// 	total := 0

// 	for _, val := range values {
// 		total += val
// 	}

// 	return total
// }
