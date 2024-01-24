package order

type OrderService interface {
	CreateOrder(id string, customerID string) (string, error)
	UpdateOrderItem(productID string, quantity int) (string, int, error)
	SubmitOrder(id string) error
}
