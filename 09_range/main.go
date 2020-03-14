package main

import "fmt"

func main() {
	ids := []int{1, 43, 12, 45, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d: %d\n", i, id)
	}

	// Ignore the index
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	emails := map[string]string{
		"test":  "test@test.com",
		"hello": "hello@hello.com",
	}

	for k, v := range emails {
		fmt.Println(k, v)
	}
}
