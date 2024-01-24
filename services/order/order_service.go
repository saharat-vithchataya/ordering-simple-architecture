package order

import domain "github.com/saharat-vithchataya/ordering/domain/order"

type orderService struct {
	orderRepo domain.OrderRepository
}

func NewOrderService(orderRepo domain.OrderRepository) domain.OrderService {
	return orderService{orderRepo: orderRepo}
}

func (srv orderService) CreateOrder(id string, customerID string) (string, error) {
	return "", nil
}
func (srv orderService) UpdateOrderItem(productID string, quantity int) (string, int, error) {
	return "", 0, nil
}
func (srv orderService) SubmitOrder(id string) error {
	return nil
}
