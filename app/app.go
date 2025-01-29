package app

import (
	"log"
	"net/http"

	"example.com/hexagonal/domain"
	"example.com/hexagonal/handlers"
	"example.com/hexagonal/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	// wiring
	//ch := handlers.CustomerHandlers{Service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := handlers.CustomerHandlers{Service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// Define the routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
