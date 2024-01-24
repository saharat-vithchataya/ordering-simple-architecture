package order

import "errors"

var (
	ErrInvalidOrder             = errors.New("invalid order")
	ErrEditAfterSubmit          = errors.New("submitted order cannot be edited")
	ErrItemQuantityLessThanZero = errors.New("order item cannot be less than 0")
	ErrDoubleSubmit             = errors.New("submitted order cannot be submitted again")
)

type Order struct {
	ID         string
	CustomerID string
	Items      map[string]int
	Submitted  bool
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
