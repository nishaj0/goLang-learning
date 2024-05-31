package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	// creating pointer
	// var ptr *int
	// fmt.Println(ptr) // <nil>

	// referencing pointer
	myNumber := 23

	var ptr = &myNumber

	fmt.Println("value of pointer", ptr)  // 0xc00000a098
	fmt.Println("value of pointer", *ptr) // 23

	// updating pointer value
	*ptr = *ptr + 2
	fmt.Println("new value", myNumber) // 25

}
