package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	age         int
	BirthDate   time.Time
}

type myStruct struct {
	FirstName string
}

func main() {
	var whatToSay string
	var saySomethingElse string
	var i int
	var world string
	var color string

	whatToSay, world = saySomething("Hello")

	log.Println(whatToSay, world)

	saySomethingElse, _ = saySomething("Goodbye")

	log.Println(saySomethingElse)
	i = 7
	log.Println(i)

	color = "Blue"
	log.Println("Color is set to", color)
	changeUsingPointer(&color)
	log.Println("Color after pointer change is set to", color)

	user := User{
		FirstName:   "Joe",
		LastName:    "Bloggs",
		PhoneNumber: "0123456789",
	}

	log.Println(user.FirstName, user.LastName)

}

func saySomething(s string) (string, string) {
	return s, "World"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	*s = "Red"
}
