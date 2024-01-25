package services

import (
	"fmt"

	domain "github.com/saharat-vithchataya/ordering/domain/order"
	"github.com/saharat-vithchataya/ordering/logs"
)

type orderService struct {
	orderRepo domain.OrderRepository
}

func NewOrderService(orderRepo domain.OrderRepository) domain.OrderService {
	return orderService{orderRepo: orderRepo}
}

func (srv orderService) CreateOrder(customerID string) (string, error) {
	orderID := srv.orderRepo.NextIdentity()

	order, err := domain.NewOrder(orderID, customerID)
	if err != nil {
		logs.Error(err)
		return "", err
	}

	err = srv.orderRepo.Save(&order)
	if err != nil {
		logs.Error(err)
		return "", err
	}

	logs.Info(fmt.Sprintf("Order created successfully. ID: %v", orderID))
	return orderID, nil
}

func (srv orderService) UpdateOrderItem(orderID string, updatedItem domain.OrderUpdateItem) error {
	order, err := srv.orderRepo.FromID(orderID)
	if err != nil {
		logs.Error(err)
		return domain.ErrOrderNotFound
	}

	if err = order.UpdateItem(updatedItem.ProductID, updatedItem.Quantity); err != nil {
		logs.Error(err)
		return err
	}

	if err = srv.orderRepo.Save(order); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (srv orderService) SubmitOrder(orderID string) error {
	order, err := srv.orderRepo.FromID(orderID)
	if err != nil {
		logs.Error(err)
		return err
	}

	if err = order.Submit(); err != nil {
		logs.Error(err)
		return err
	}

	if err = srv.orderRepo.Save(order); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (srv orderService) GetOrder(orderID string) (*domain.OrderResponse, error) {
	order, err := srv.orderRepo.FromID(orderID)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &domain.OrderResponse{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Items:      order.Items,
		Submitted:  order.Submitted,
	}, nil
}
