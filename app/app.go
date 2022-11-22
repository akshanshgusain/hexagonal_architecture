package app

import (
	"log"
	"net/http"
)

func Start() {
	// routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// Starting Server
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
