package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")

	john := User{"John", "john@gmail.com", true, 23}

	fmt.Println(john) // {John john@gmail.com true 23}

	fmt.Printf("%v\n", john)                                  // {John john@gmail.com true 23}
	fmt.Printf("%+v\n", john)                                 // {Name:John Email:john@gmail.com Status:true Age:23}
	fmt.Printf("name is: %+v age is: %v", john.Name, john.Age) // name is: John age is: 23
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
