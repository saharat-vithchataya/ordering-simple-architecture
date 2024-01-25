package customer

type CustomerService interface {
	CreateNewCustomer(CustomerRequest) (string, error)
	GetCustomer(string) (*CustomerResponse, error)
}

type CustomerResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type CustomerRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
