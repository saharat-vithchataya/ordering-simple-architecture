package customer

import (
	"fmt"

	domain "github.com/saharat-vithchataya/ordering/domain/customer"
	"github.com/saharat-vithchataya/ordering/logs"
)

type customerService struct {
	custRepo domain.CustomerRepository
}

func NewCustomerService(custRepo domain.CustomerRepository) domain.CustomerService {
	return customerService{custRepo: custRepo}
}

func (srv customerService) CreateNewCustomer(name string) (string, error) {
	id := srv.custRepo.NextIdentity()

	newCustomer, err := domain.NewCustomer(id, name)
	if err != nil {
		logs.Error(err)
		return "", err
	}

	if err = srv.custRepo.Save(newCustomer); err != nil {
		logs.Error(err)
		return "", err
	}

	logs.Info(fmt.Sprintf("customer %v sfully created a new account.", id))
	return id, nil
}

func (srv customerService) GetCustomer(id string) (domain.Customer, error) {
	customer, err := srv.custRepo.FromID(id)
	if err != nil {
		logs.Error(err)
		return domain.Customer{}, err
	}

	return customer, nil
}
