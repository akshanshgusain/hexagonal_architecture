package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// Server side PORT/ Secondary PORT

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, error)
}
