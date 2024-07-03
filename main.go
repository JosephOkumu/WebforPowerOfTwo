package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// isPowerOfTwo checks if a given number is a power of two
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

// checkHandler is an HTTP request handler that checks if a number is a power of two
func checkHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "number" parameter from the URL query string
	numberStr := r.FormValue("number")

	// Convert the string representation of the number to an integer
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		// If the conversion fails, send a 400 Bad Request response
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	// Check if the number is a power of two
	result := isPowerOfTwo(number)

	// Create an anonymous struct to hold the number and result
	data := struct {
		Number int
		Result bool
	}{
		Number: number,
		Result: result,
	}

	// Parse the "result.html" template file
	tmpl := template.Must(template.ParseFiles("templates/result.html"))

	// Execute the template with the data and write the response
	tmpl.Execute(w, data)
}

func main() {
	// Register the "/check" endpoint with the checkHandler
	http.HandleFunc("/check", checkHandler)

	// Register the root "/" endpoint to serve the index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the "index.html" template file
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		// Execute the template and write the response
		tmpl.Execute(w, nil)
	})

	// Print a message indicating that the server has started
	fmt.Println("Server started at :8080")

	// Start the HTTP server and listen on port 8080
	http.ListenAndServe(":8080", nil)
}
