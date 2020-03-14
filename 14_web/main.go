package main

import (
	"fmt"
	"net/http"
)

// Takes two parameters:
// http.ResponseWriter
// *http:Request <- Note that this is a pointer
func index(w http.ResponseWriter, r *http.Request) {
	// Raw string
	var html string = `
	<!doctype HTML>
	<head>
		<title>Home page</title>
	</head>
	<body>
		<h1>Header</h1>
		<p>Bismillahir Rahmanir Rahim, ALLAHU AKBAR</p>
		<p>Alhamdulillah</p>
		<p>Hello World</p>
	</body>
	</html>
	`

	// Write the string into the ResponseWriter
	fmt.Fprintf(w, html)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey this is the about page!")
}

func main() {
	// Register the different URL patterns (similar to routes)
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	addr := "localhost:3000"
	fmt.Printf("Starting server at %s...\n", addr)

	// Start the server and listen on a port
	http.ListenAndServe(addr, nil)
}
