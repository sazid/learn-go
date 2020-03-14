package main

import "fmt"

/*
	Functions are defined in the following format:
	func funcName(param1 type, param2 type) returnType {
		...
		return xyz
	}
*/
func greeting(name string) string {
	return "Hello " + name
}

func getSum(a int, b int) int {
	return a + b
}

/*
	When there are multiple consecutive parameters of the same type,
	the type can be omitted for all of them except the last one.
	In this case, both the parameters a and b are of type int
*/
func getSumShortParamType(a, b int, c, d float64) float64 {
	return float64(a+b) + c + d
}

/*
	A function can return multiple values, the corresponding return
	type must match the types of the returned values
*/
func mySwap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(greeting("sazid"))
	fmt.Println(getSum(2, 3))
	fmt.Println(getSumShortParamType(2, 3, 2.5, 32.5312))

	a, b := 1, 5
	a, b = mySwap(a, b)
	fmt.Println(a, b)

	// When any return value is not needed, it can be discarded with
	// the blank identifier _.
	_, c := mySwap(10, 20)
	fmt.Println(c)
}
