package services

import (
	domain "github.com/saharat-vithchataya/ordering/domain/product"
	"github.com/saharat-vithchataya/ordering/logs"
)

type productService struct {
	prodRepo domain.ProductRepository
}

func NewProductService(prodRepo domain.ProductRepository) domain.ProductService {
	return productService{prodRepo: prodRepo}
}

func (srv productService) CreateNewProduct(request domain.ProductCreate) (string, error) {
	productID := srv.prodRepo.NextIdentity()
	product, err := domain.NewProduct(productID, request.Name, request.Price, request.Quantity)
	if err != nil {
		logs.Error(err)
		return "", err
	}

	srv.prodRepo.Save(&product)

	return productID, nil
}
func (srv productService) GetProduct(productID string) (*domain.ProductResponse, error) {
	product, err := srv.prodRepo.FromID(productID)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &domain.ProductResponse{
		ID:       productID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}
