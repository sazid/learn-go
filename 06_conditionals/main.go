package main

import "fmt"

func main() {
	x := 10
	y := 10

	/*
		Conditions can be written without any parenthesis in an
		if statement

		if condition {

		} else if condition2 {

		} else {

		}
	*/
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equals to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	/*
		Switch is quite powerful in go. They even allow conditions
		in cases instead of just constants.

		Switch can be used instead of if else statements by omitting
		the variable to switch on

		Also notice that, there is no need for break statements!
	*/
	switch 'c' {
	case 'c':
		fmt.Println("Hey its C!")
	case 'b':
		fmt.Println("Its B!")
	}

	num := 3
	switch {
	case num < 10:
		fmt.Println("num is less than 10")
	case num > 10:
		fmt.Println("num is greater than 10")
	case num == 10:
		fmt.Println("num is equal to 10")
	default:
		fmt.Println("Impossible!")
	}
}
