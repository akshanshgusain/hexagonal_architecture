package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customers struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			log.Fatal("XML Encode Error")
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			log.Fatal("JSON Encode Error")
		}
	}
}
