package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.request) {
	fmt.Fprintf(w, "Alfian Nova Handhika")
}

func main() {
	http.HandhleFunc("/", helloWorld)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}