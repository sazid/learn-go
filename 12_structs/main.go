package main

import (
	"fmt"
	"strconv"
)

// Person : Store a person's information
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// greet : Greeting method (value reciever) for Person struct
// Notice that its defined outside the struct, more like an extension
// methods in other languages like kotlin, c#
func (p Person) greet() string {
	var s string = "Hello, " + p.firstName + " " + p.lastName +
		" is " + strconv.Itoa(p.age) + " years old."

	// Note that, changing the age here doesn't alter the actual
	// object's age. Everything in go is "pass by value" by default.
	// You'll have to use pointer recievers to get the reference to
	// the actual object.
	p.age = 20
	return s
}

// hasBirthday : hasBirthday method (pointer reciever), modifies
// the struct itself. Notice that, there's no special syntax for
// accessing a pointer reciever, you use it like a value reciever
// This is similar to pass by reference in c++
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried : pointer reciever
func (p *Person) getMarried(spouse *Person) {
	fmt.Println("Congratulations on getting married with", spouse.firstName)
	if p.gender == "m" {
		fmt.Println("Enjoy the new life with your wife!")
	} else if p.gender == "f" {
		fmt.Println("Enjoy the new life with your husband!")
	}
}

func main() {
	/*
		There are no classes in go. Only structs are used
		to represent data
	*/

	// Init person using struct
	person1 := Person{
		firstName: "Bob",
		lastName:  "Pike",
		city:      "Nowhere",
		gender:    "m",
		age:       30,
	}

	// Alternative
	person2 := Person{}
	person2.firstName = "Black"
	person2.lastName = "White"
	person2.city = "Space"
	person2.gender = "f"
	person2.age = 43

	fmt.Println(person1, person2)

	// Access the methods as if they are part of the Person struct
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried(&person2)
}
