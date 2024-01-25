package order

type OrderService interface {
	CreateOrder(customerID string) (string, error)
	UpdateOrderItem(orderID string, updatedItem OrderUpdateItem) error
	GetOrder(orderID string) (*OrderResponse, error)
	SubmitOrder(orderID string) error
}

type OrderResponse struct {
	ID         string         `json:"id"`
	CustomerID string         `json:"customer_id"`
	Items      map[string]int `json:"items"`
	Submitted  bool           `json:"submitted"`
}

type OrderUpdateItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
