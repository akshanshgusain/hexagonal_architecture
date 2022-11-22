package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customers struct {
	Name    string
	City    string
	Zipcode string
}

func main() {
	// routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// Starting Server
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Handlers

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello Welcome to the Hexagonal Architecture APP")
	if err != nil {
		return
	}
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{"Akshansh Gusain", "New Delhi", "110001"},
		{"Priyanka Khurana", "New Delhi", "110001"},
		{"Rachit Kawar", "Jodhpur", "132901"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
