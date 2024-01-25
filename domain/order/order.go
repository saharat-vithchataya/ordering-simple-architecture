package order

import (
	"errors"
)

var (
	ErrInvalidOrder             = errors.New("invalid order")
	ErrEditAfterSubmit          = errors.New("submitted order cannot be edited")
	ErrItemQuantityLessThanZero = errors.New("order item cannot be less than 0")
	ErrDoubleSubmit             = errors.New("submitted order cannot be submitted again")
	ErrOrderNotFound            = errors.New("order not found")
)

type Order struct {
	ID         string         `bson:"_id"`
	CustomerID string         `bson:"customer_id"`
	Items      map[string]int `bson:"items"`
	Submitted  bool           `bson:"submitted"`
}

func NewOrder(id string, customerID string) (Order, error) {
	if id == "" || customerID == "" {
		return Order{}, ErrInvalidOrder
	}
	return Order{
		ID:         id,
		CustomerID: customerID,
		Items:      map[string]int{},
		Submitted:  false,
	}, nil
}

func (entity *Order) UpdateItem(productID string, quantity int) error {
	if entity.Submitted {
		return ErrEditAfterSubmit
	}

	if quantity < 0 {
		return ErrItemQuantityLessThanZero
	}

	if quantity == 0 {
		delete(entity.Items, productID)
		return nil
	}

	entity.Items[productID] = quantity

	return nil
}

func (entity *Order) Submit() error {
	if entity.Submitted {
		return ErrDoubleSubmit
	}

	entity.Submitted = true

	return nil
}
