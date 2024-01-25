package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/saharat-vithchataya/ordering/domain/product"
)

type productHandler struct {
	prodSrv domain.ProductService
}

func NewProductHandler(prodSrv domain.ProductService) productHandler {
	return productHandler{prodSrv: prodSrv}
}

func (h productHandler) CreateNewProduct(c *fiber.Ctx) error {
	var product domain.ProductCreate
	if err := c.BodyParser(&product); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	productID, err := h.prodSrv.CreateNewProduct(product)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"product_id": productID,
	})
}

func (h productHandler) GetProduct(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	response, err := h.prodSrv.GetProduct(productID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
