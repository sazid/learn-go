package main

import "fmt"

func main() {
	/*
		There is only one loop construct in go and that is the for.
		However it can have three different forms.
	*/

	// More like while(...) in other languages
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
}
