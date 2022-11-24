package app

import (
	"fmt"
	"github.com/akshanshgusain/Hexagonal-Architecture/domain"
	"github.com/akshanshgusain/Hexagonal-Architecture/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// Sanity test is a test of newly deployed environment.
// responsible to check if all the env variables are passed correctly
func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("environment variables are empty")
	}
}

func Start() {
	//Sanity Check
	sanityCheck()

	// Create a multiplexer
	router := mux.NewRouter()

	// wiring
	// handler -> Service -> repository (dependency injection)
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// routes

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// Starting Server
	// OS env: SERVER_ADDRESS=localhost SERVER_PORT=8080 go run main.go
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
