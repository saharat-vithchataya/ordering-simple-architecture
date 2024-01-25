package product

type ProductService interface {
	CreateNewProduct(ProductCreate) (string, error)
	GetProduct(id string) (*ProductResponse, error)
}

type ProductCreate struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"int"`
}

type ProductResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
