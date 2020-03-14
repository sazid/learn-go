package main

import (
	"fmt"
	"math"

	"github.com/sazid/go_crash_course/03_packages/strutil"
)

/*
	Multiple packages are imported with the
	import (
		"package1"
		"package2"
		...
	)
	syntax
*/

func main() {
	fmt.Println(math.Floor(2.7), math.Ceil(2.7))
	fmt.Println(math.Pi)
	fmt.Printf("%T\n", math.Pi)

	/*
		Imported packages should be prefixed before accessing any of the
		packages' functions and other stuffs
	*/
	reversedName := strutil.Reverse("Sazid")
	fmt.Println(reversedName)
}
