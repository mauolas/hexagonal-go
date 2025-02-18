package dto

type CustomerResponse struct {
	Id          string `json:"id" xml:"id" `
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city" `
	Zipcode     string `json:"zipcode" xml:"zipcode" `
	DateofBirth string `json:"date_of_birth" xml:"dateofbirth" `
	Status      string `json:"status" xml:"status"`
}
