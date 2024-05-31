package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println(fruitList)      //[Apple Orange  Banana]
	fmt.Println(len(fruitList)) //4
}
