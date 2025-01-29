package domain

import "example.com/hexagonal/errs"

// Customer struct
type Customer struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateofBirth string `json:"dateofbirth" xml:"dateofbirth"`
	Status      string `json:"status" xml:"status"`
}

// CustomerRepository interface
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
