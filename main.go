package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	numberStr := r.FormValue("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	result := isPowerOfTwo(number)
	data := struct {
		Number int
		Result bool
	}{
		Number: number,
		Result: result,
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
