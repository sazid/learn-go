package main

import "fmt"

func main() {
	// Array
	var fruitArr [3]string

	// Assign values to array
	fruitArr[0] = "Apple"
	fruitArr[1] = "Grape"
	fruitArr[2] = "Mango"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[2])

	// Declare and assign
	// Notice that, only a few of the items are assigned
	// Both the length and type are part of the "type" of the array.
	// In this case, [10]int is the type of the array
	fiboArr := [10]int{0, 1, 1, 2, 3, 5}
	fmt.Println(fiboArr)
	fmt.Printf("%T\n", fiboArr)

	// Get the length using len() function
	fmt.Println(len(fiboArr))

	// Slices
	fruitSlice := []string{"apple", "grape", "mango"}
	fmt.Println(fruitSlice)
	fmt.Printf("%T\n", fruitSlice)

	// append() returns a new slice
	fruitSlice = append(fruitSlice, "banana")
	fruitSlice = append(fruitSlice, "pineapple")

	fmt.Println(fruitSlice, len(fruitSlice))
	fmt.Println(fruitSlice[1:4])

	// Slices can be made using the make(type, length) function
	newSlice := make([]int, 3)
	newSlice[1] = 5
	fmt.Println(newSlice)
}
