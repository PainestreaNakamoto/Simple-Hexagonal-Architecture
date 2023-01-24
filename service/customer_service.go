package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	RepoCustomer repository.CustomerRepository
}

func NewCustomerService(RepoCustomer repository.CustomerRepository) CustomerService {
	return customerService{RepoCustomer: RepoCustomer}
}

func (ser customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := ser.RepoCustomer.GetAll()

	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	customer_responses := []CustomerResponse{}
	for _, customer := range customers {
		customer_response := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customer_responses = append(customer_responses, customer_response)
	}

	return customer_responses, nil
}

func (ser customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := ser.RepoCustomer.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customer_response := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &customer_response, nil
}
