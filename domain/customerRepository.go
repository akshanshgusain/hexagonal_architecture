package domain

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/dto"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
)

// Server side PORT/ Secondary PORT

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c *Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

/*domain has the complete knowledge of the DTO thus, we should be able to ask the domain to get the DTO representation*/

func (c *Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}
