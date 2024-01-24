package product

type ProductService interface {
	CreateNewProduct(name string, price float64, quantity int) (string, error)
	GetProduct(id string) (Product, error)
}
