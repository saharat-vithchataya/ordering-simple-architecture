package order

import (
	domain "github.com/saharat-vithchataya/ordering/domain/order"
)

type orderRepositoryMongoDB struct {
	// mongodb driver
	collection string
}

func NewOrderRepositoryMongoDB(collection string) domain.OrderRepository {
	return orderRepositoryMongoDB{collection: collection}
}

func (repo orderRepositoryMongoDB) NextIdentity() string {
	return "testing from mongodb"
}

func (repo orderRepositoryMongoDB) FromID(string) (domain.Order, error) {
	return domain.Order{}, nil
}

func (repo orderRepositoryMongoDB) Save(domain.Order) error {
	return nil
}
