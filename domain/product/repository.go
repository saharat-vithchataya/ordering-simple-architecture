package product

type ProductRepository interface {
	NextIdentity() string
	FromID(string) (*Product, error)
	Save(*Product) error
}
