package customer

type CustomerRepository interface {
	NextIdentity() string
	FromID(string) (*Customer, error)
	Save(*Customer) error
}
