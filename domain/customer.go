package domain

import (
	"example.com/hexagonal/dto"
	"example.com/hexagonal/errs"
)

// Customer struct
type Customer struct {
	Id          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"name" xml:"name" db:"name"`
	City        string `json:"city" xml:"city" db:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode" db:"zipcode"`
	DateofBirth string `json:"dateofbirth" xml:"dateofbirth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status" db:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

// CustomerRepository interface
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
	FindByStatus(string) ([]Customer, *errs.AppError)
}
