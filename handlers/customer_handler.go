package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/saharat-vithchataya/ordering/domain/customer"
)

type customerHandler struct {
	custSrv domain.CustomerService
}

func NewCustomerHandler(custSrv domain.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

func (h customerHandler) CreateNewCustomer(c *fiber.Ctx) error {
	var entity domain.CustomerRequest
	c.BodyParser(&entity)
	customerID, err := h.custSrv.CreateNewCustomer(entity)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"customer_id": customerID,
	})
}

func (h customerHandler) GetCustomer(c *fiber.Ctx) error {
	customerID := c.Params("customer_id")
	customer, err := h.custSrv.GetCustomer(customerID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(customer)
}
