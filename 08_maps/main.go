package main

import "fmt"

func main() {
	// Create map using map[keyType]valueType syntax
	// emails := make(map[string]string)
	emails := map[string]string{"tray": "tray@tray.com"}

	// Assign key value
	emails["test"] = "test@test.com"
	emails["hello"] = "hello@hello.com"
	emails["bob"] = "bob@bob.com"

	// Access a map
	fmt.Println(emails, len(emails))
	fmt.Println(emails["test"])

	// Delete from map
	delete(emails, "hello")
	fmt.Println(emails, len(emails))
}
