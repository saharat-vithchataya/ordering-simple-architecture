package customer

type CustomerService interface {
	CreateNewCustomer(string) (string, error)
	GetCustomer(string) (Customer, error)
}
