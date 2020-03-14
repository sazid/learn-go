package main

/*
	package main - indicates that this package will be used as an executable
	rather than as a shared library.
*/

import "fmt"

/*
	Import the "fmt" package. All the files in a package are visible to one
	another, i.e variables, functions, etc., can be accessed from any file
	within the same package. This simplifies how different types are visible
	to one another compared to other languages like java, c# where you have
	to explicitly tell or have to remember the default visibility modifier.
*/

/*
	Functions are declared with the 'func' keyword, followed by the name of
	the function. func main() { ... }

	main() function is the entry point for the executable program,
	i.e this function will be called automatically when the program executes
*/
func main() {
	fmt.Println("Hello World")
}
