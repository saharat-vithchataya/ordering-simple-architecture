package customer

import "errors"

var (
	ErrInvalidCustomer  = errors.New("invalid customer information")
	ErrCustomerNotFound = errors.New("customer not found")
)

type Customer struct {
	ID      string `bson:"_id"`
	Name    string `bson:"name"`
	Phone   string `bson:"phone"`
	Address string `bson:"address"`
}

func NewCustomer(id, name string) (Customer, error) {
	if name == "" || id == "" {
		return Customer{}, ErrInvalidCustomer
	}
	return Customer{ID: id, Name: name}, nil
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
