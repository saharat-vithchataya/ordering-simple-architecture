package product

import domain "github.com/saharat-vithchataya/ordering/domain/product"

type productRepositoryMongoDB struct {
	// mongodb driver
	collection string
}

func NewProductRepositoryMongoDB(collection string) domain.ProductRepository {
	return productRepositoryMongoDB{collection: collection}
}

func (repo productRepositoryMongoDB) NextIdentity() string {
	return "testing from mongodb"
}

func (repo productRepositoryMongoDB) FromID(string) (domain.Product, error) {
	return domain.Product{}, nil
}

func (repo productRepositoryMongoDB) Save(domain.Product) error {
	return nil
}
