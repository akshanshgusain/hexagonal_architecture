package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	// Create a multiplexer
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// Starting Server
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
