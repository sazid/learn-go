package main

import "fmt"

// Declaring a global variable, see below
var myGlobalVar = 10

// This will produce error, see below
// myShortGlobalVar := 10

func main() {
	/*
		MAIN TYPES

		string
		bool
		// int32 used in 32-bit and int64 in 64-bit systems
		int int8 int16 int32 int64
		uint uint8 uint16 unit32 uint64 uintptr
		byte - alias for uint8
		rune - alias for int32 - Use this to guarantee int32
		flaot32 float64(default)
		complex64 complex128
	*/

	/*
		Using var & const
		var variableName [type] = value
		const variableName [type] = value

		[type] in square brackets indicate that this is optional and most of
		the time, the compiler should be able to infer the types.

		Note: go convention is to use camelCase for variable names
	*/
	var name string = "sazid"
	var inferType = "test"
	var age int8 = 127
	const isCool = true
	// isCool = false // This will produce error since isCool is const

	/*
		'var' MUST be used if declared,
		however,
		'const' can be declared without being used anywhere

		Try declaring a var without using it anywhere and compile the file.
		Try the same for a const.
	*/

	fmt.Println(name, age)
	fmt.Println(inferType)
	// %T to print type of something
	fmt.Printf("%T %T\n", age, inferType)

	/*
		Using short-hand notation for variable declaration
		varName := value

		Limitation: short-hand declaration cannot be used for declaring global
		variables.
	*/
	shortVar := "xyz"
	abc, def := 10, 20

	fmt.Println(shortVar, abc, def)
	fmt.Printf("%T", shortVar)
}
