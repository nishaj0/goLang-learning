package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")

	john := User{"John", "john@gmail.com", true, 23}

	// fmt.Println(john) // {John john@gmail.com true 23}

	// fmt.Printf("%v\n", john)                                   // {John john@gmail.com true 23}
	// fmt.Printf("%+v\n", john)                                  // {Name:John Email:john@gmail.com Status:true Age:23}
	// fmt.Printf("name is: %+v age is: %v", john.Name, john.Age) // name is: John age is: 23

	john.GetStatus()
	john.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is he active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "new@email"
	fmt.Println(u.Email)
}
