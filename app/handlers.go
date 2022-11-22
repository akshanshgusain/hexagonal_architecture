package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/akshanshgusain/Hexagonal-Architecture/service"
	"log"
	"net/http"
)

type Customers struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

// Handlers

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customers{
	//	{"Akshansh Gusain", "New Delhi", "110001"},
	//	{"Priyanka Khurana", "New Delhi", "110001"},
	//	{"Rachit Kawar", "Jodhpur", "132901"},
	//}
	customers, _ := ch.service.GetAllCustomers()

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
