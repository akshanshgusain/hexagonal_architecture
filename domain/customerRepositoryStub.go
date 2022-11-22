package domain

// CustomerRepositoryStub Stub
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Foo", "Dublin", "000000", "10-Nov-1987", "1"},
		{"2", "Bar", "Tokyo", "000000", "10-Nov-1987", "1"},
		{"3", "Baz", "Paris", "000000", "10-Nov-1987", "1"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
