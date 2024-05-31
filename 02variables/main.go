package main

import "fmt"

const LoginToken string = "ffdskfjasdkfjsdj"

func main() {
	var username string = "john"
	fmt.Println(username)
	fmt.Printf("variable type is %T \n", username)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("variable type is %T \n", isLogged)
	fmt.Printf("variable type is %T \n", username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable type is %T \n", smallVal)

	// ? default value
	var anotherVar int
	fmt.Println(anotherVar)                          // 0
	fmt.Printf("variable type is %T \n", anotherVar) // variable type is int

	var anotherVar1 string
	fmt.Println(anotherVar1)                          // ''
	fmt.Printf("variable type is %T \n", anotherVar1) // variable type is string

	// ? implicit type
	var website = "google.com"
	fmt.Println(website)                          // google.com
	fmt.Printf("variable type is %T \n", website) // variable type is string

	numberOfUsers := 30000.0
	fmt.Println(numberOfUsers) // 30000.0

	fmt.Println(LoginToken)                          // ffdskfjasdkfjsdj
	fmt.Printf("variable type is %T \n", LoginToken) // variable type is string
}
