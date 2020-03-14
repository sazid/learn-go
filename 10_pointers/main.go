package main

import "fmt"

func main() {
	a := 5
	// Point to memory address of a
	// b := &a
	var b *int = &a

	// Change the value at the pointed address
	*b = 12

	// Access the value at the pointed address
	fmt.Println(a, *b)
}
