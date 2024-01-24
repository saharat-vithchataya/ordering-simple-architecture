package customer

import (
	domain "github.com/saharat-vithchataya/ordering/domain/customer"
)

type customerRepositoryMongoDB struct {
	// mongodb driver
	collection string
}

func NewCustomerRepositoryMongoDB(collection string) domain.CustomerRepository {
	return customerRepositoryMongoDB{collection: collection}
}

func (repo customerRepositoryMongoDB) NextIdentity() string {
	return "testing from mongodb"
}

func (repo customerRepositoryMongoDB) FromID(string) (domain.Customer, error) {
	return domain.Customer{}, nil
}

func (repo customerRepositoryMongoDB) Save(domain.Customer) error {
	return nil
}
