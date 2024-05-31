package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, value := range days {
	// 	fmt.Printf("index: %v, its value: %v\n", index, value)
	// }

	val := 1
	for val < 5 {

		if val == 3 {
			goto hm
		}

		fmt.Println("value is", val)
		val++
	}

hm:
	fmt.Println("jumping at hello world")
}
