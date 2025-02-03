package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"example.com/hexagonal/errs"
	"example.com/hexagonal/logger"
	"example.com/hexagonal/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.Service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.Service.GetCustomer(id)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage(), "application/json")
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		WriteResponse(w, http.StatusOK, customer, "application/xml")
	} else {
		WriteResponse(w, http.StatusOK, customer, "application/json")
	}

}

func (ch *CustomerHandlers) GetCustomerByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	if status != "active" && status != "inactive" {
		err := errs.NewBadRequestError("status should be either active or inactive")
		logger.Error("Error while validating status: " + err.Message)
		WriteResponse(w, err.Code, err.AsMessage(), "application/json")
		return
	}

	customers, err := ch.Service.GetCustomerByStatus(status)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage(), "application/json")
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		WriteResponse(w, http.StatusOK, customers, "application/xml")
	} else {
		WriteResponse(w, http.StatusOK, customers, "application/json")
	}
}

func WriteResponse(w http.ResponseWriter, code int, data interface{}, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
	if contentType == "application/xml" {
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	} else {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}
