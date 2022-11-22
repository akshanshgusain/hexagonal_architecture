package app

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/domain"
	"github.com/akshanshgusain/Hexagonal-Architecture/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	// Create a multiplexer
	router := mux.NewRouter()

	// wiring
	// handler -> Service -> repository (dependency injection)
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// routes

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// Starting Server
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
