package product

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

func (srv productService) CreateNewProduct(name string, price float64, quantity int) (string, error) {
	productID := srv.prodRepo.NextIdentity()
	if _, err := domain.NewProduct(productID, name, price, quantity); err != nil {
		logs.Error(err)
		return "", err
	}

	return productID, nil
}
func (srv productService) GetProduct(productID string) (domain.Product, error) {
	product, err := srv.prodRepo.FromID(productID)
	if err != nil {
		logs.Error(err)
		return domain.Product{}, err
	}

	return product, nil
}
