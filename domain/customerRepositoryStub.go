package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "John Doe", "New York", "10001", "1990-01-01", "1"},
		{"1002", "Jane Doe", "New York", "10001", "1990-01-01", "1"},
	}

	return CustomerRepositoryStub{customers}
}
