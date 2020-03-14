package main

import "fmt"

// This function returns another function with the signature
// func(int) -> int
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// The adder() function returns another function
	// We store that function in sum
	sum := adder()

	// Type of sum is a function f:int -> int
	fmt.Printf("%T\n", sum)

	for i := 1; i <= 10; i++ {
		fmt.Println(sum(i))
	}

	// Anonymous functions can be excuted on spot using this syntax
	// func(param1, param2, ...) {
	//		...
	//  }(arg1, arg2, ...)
	x := 10
	func(printVar int) {
		fmt.Println(printVar)
	}(x)
}
