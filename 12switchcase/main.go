package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("welcome to switchCase")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("dice number", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("start game")

	case 2:
		fmt.Println("move 2 spaces")

	case 3:
		fmt.Println("move 3 spaces")
		fallthrough
	case 4:
		fmt.Println("move 4 spaces")

	case 5:
		fmt.Println("move 5 spaces")

	case 6:
		fmt.Println("move 6 spaces and roll dice again")
	default:
		fmt.Println("what was that")

	}
}
