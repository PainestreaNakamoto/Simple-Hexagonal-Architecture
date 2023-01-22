package service

import "bank/repository"

type customerService struct {
	RepoCustomer repository.CustomerRepository
}

func NewCustomerService(RepoCustomer repository.CustomerRepository) CustomerService {
	return customerService{RepoCustomer: RepoCustomer}
}

func (ser customerService) GetCustomers() ([]CustomerResponse, error) {
	return nil, nil
}

func (ser customerService) GetCustomer(id int) (*CustomerResponse, error) {
	return nil, nil
}
