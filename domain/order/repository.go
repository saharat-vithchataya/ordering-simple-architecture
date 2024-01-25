package order

type OrderRepository interface {
	NextIdentity() string
	FromID(string) (*Order, error)
	Save(*Order) error
}
