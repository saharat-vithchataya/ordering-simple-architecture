package product

import "errors"

var (
	ErrInvalidProduct         = errors.New("invalid product")
	ErrPriceIsLessThanZero    = errors.New("price is less than zero")
	ErrQuantityIsLessThanZero = errors.New("quantity is less than zero")
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func NewProduct(id, name string, price float64, quantity int) (Product, error) {
	if id == "" || name == "" {
		return Product{}, ErrInvalidProduct
	}

	if price < 0 {
		return Product{}, ErrPriceIsLessThanZero
	}

	if quantity < 0 {
		return Product{}, ErrPriceIsLessThanZero
	}

	return Product{
		ID:          id,
		Name:        name,
		Description: "",
		Price:       price,
		Quantity:    quantity,
	}, nil
}

func (entity *Product) UpdatePrice(newPrice float64) error {
	if newPrice < 0 {
		return ErrPriceIsLessThanZero
	}
	entity.Price = newPrice

	return nil
}

func (entity *Product) UpdateQuantity(qty int) error {
	if qty < 0 {
		return ErrQuantityIsLessThanZero
	}
	entity.Quantity = qty

	return nil
}
