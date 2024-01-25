package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/saharat-vithchataya/ordering/domain/order"
)

type orderHandler struct {
	orderSrv domain.OrderService
}

func NewOrderHandler(orderSrv domain.OrderService) orderHandler {
	return orderHandler{orderSrv: orderSrv}
}

func (h orderHandler) CreateNewOrder(c *fiber.Ctx) error {
	customerID := c.Params("customer_id")
	orderID, err := h.orderSrv.CreateOrder(customerID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"order_id": orderID,
	})
}

func (h orderHandler) GetOrder(c *fiber.Ctx) error {
	orderID := c.Params("order_id")
	order, err := h.orderSrv.GetOrder(orderID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (h orderHandler) UpdateOrderItem(c *fiber.Ctx) error {
	orderID := c.Params("order_id")
	var updatedItem domain.OrderUpdateItem
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if err := h.orderSrv.UpdateOrderItem(orderID, updatedItem); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h orderHandler) SubmitOrder(c *fiber.Ctx) error {
	orderID := c.Params("order_id")
	if err := h.orderSrv.SubmitOrder(orderID); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
