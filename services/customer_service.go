package services

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

func (srv customerService) CreateNewCustomer(request domain.CustomerRequest) (string, error) {
	id := srv.custRepo.NextIdentity()

	newCustomer, err := domain.NewCustomer(id, request.Name)
	newCustomer.Phone = request.Phone
	newCustomer.Address = request.Address

	if err != nil {
		logs.Error(err)
		return "", err
	}

	if err = srv.custRepo.Save(&newCustomer); err != nil {
		logs.Error(err)
		return "", err
	}

	logs.Info(fmt.Sprintf("customer %v sfully created a new account.", id))
	return id, nil
}

func (srv customerService) GetCustomer(id string) (*domain.CustomerResponse, error) {
	customer, err := srv.custRepo.FromID(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &domain.CustomerResponse{
		ID:      customer.ID,
		Name:    customer.Name,
		Phone:   customer.Phone,
		Address: customer.Address,
	}, nil
}
