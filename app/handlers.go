package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/akshanshgusain/Hexagonal-Architecture/service"
	"github.com/gorilla/mux"
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

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cId := vars["customer_id"]
	c, err := ch.service.GetCustomer(cId)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(c)
		if err != nil {
			log.Println("getCustomer: error encoding customer")
			return
		}
	}
}
