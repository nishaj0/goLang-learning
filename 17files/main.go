package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files")
	content := "this needs to in a file"

	file, err := os.Create("./mygofile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is", length)

	file.Close()

	readFile("./mygofile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)

	fmt.Println("text data inside the file: ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// * o/p

// welcome to files
// length is 23
// text data inside the file [116 104 105 115 32 110 101 101 100 115 32 116 111 32 105 110 32 97 32 102 105 108 101]
