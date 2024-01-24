package customer

import "errors"

var (
	ErrInvalidCustomer = errors.New("invalid customer information")
)

type Customer struct {
	ID      string
	Name    string
	Phone   string
	Address string
}

func NewCustomer(id, name string) (Customer, error) {
	if name == "" || id == "" {
		return Customer{}, ErrInvalidCustomer
	}
	return Customer{ID: id}, nil
}

func (entity *Customer) UpdateName(newName string) error {
	if newName == "" {
		return ErrInvalidCustomer
	}

	entity.Name = newName

	return nil
}

func (entity *Customer) UpdatePhone(newPhone string) error {
	if newPhone == "" {
		return ErrInvalidCustomer
	}

	entity.Phone = newPhone

	return nil
}

func (entity *Customer) UpdateAddress(newAddress string) error {
	if newAddress == "" {
		return ErrInvalidCustomer
	}

	entity.Address = newAddress

	return nil
}
