package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to slices")

	// var fruitList = []string{"Apple", "Potato", "Banana"}

	// fmt.Println(fruitList)                           // [Apple Potato Banana]
	// fmt.Printf("type of fruitList %T \n", fruitList) // type of fruitList []string

	// fruitList = append(fruitList, "Mango", "Tomato") // [Apple Potato Banana Mango Tomato]
	// fmt.Println(fruitList)

	// fmt.Println(fruitList[1:])  // [Potato Banana Mango Tomato] => fruitList[1] to len(fruitList)
	// fmt.Println(fruitList[1:3]) // [Potato Banana] => fruitList[1] to fruitList[3]

	// fruitList = append(fruitList[:3]) // [Apple Potato Banana]
	// fmt.Println(fruitList)

	// slices using make()
	// highScores := make([]int, 4)

	// highScores[0] = 392
	// highScores[1] = 924
	// highScores[2] = 491
	// highScores[3] = 335
	// // highScores[4] = 532

	// highScores = append(highScores, 555, 243, 346)
	// sort.Ints(highScores)

	// fmt.Println(highScores)                     // [243 335 346 392 491 555 924]
	// fmt.Println(sort.IntsAreSorted(highScores)) // true

	var courses = []string{"react.js", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses) // [react.js javascript swift python ruby]

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) // [react.js javascript python ruby]
}
