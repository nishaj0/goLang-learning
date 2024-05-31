package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("welcome to pizza app")
	fmt.Println("please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // (stringValue, bitSize)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to numRating", numRating+1)
	}
}
