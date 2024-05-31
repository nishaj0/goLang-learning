package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // read standard input
	fmt.Println("enter the rating of our pizza:")

	// comma ok || comma err
	input, _ := reader.ReadString('\n') // here we don't care about err
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of rating %T", input)
}
