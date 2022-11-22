package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handlers
	http.HandleFunc("/greet", greet)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello Welcome to the Hexagonal Architecture APP")
	if err != nil {
		return
	}
}
