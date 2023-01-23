package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func InitializeCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1001, Name: "Aidjf", City: "New dieai", ZipCode: "1232", Status: 1},
		{CustomerID: 2232, Name: "diji", City: "New dieai", ZipCode: "1232", Status: 2},
	}
	return customerRepositoryMock{customers: customers}
}

func (repo customerRepositoryMock) GetAll() ([]Customer, error) {
	return repo.customers, nil
}

func (repo customerRepositoryMock) GetByID(id int) (*Customer, error) {
	for _, item := range repo.customers {
		if item.CustomerID == id {
			return &item, nil
		}
	}
	return nil, errors.New("Not found")
}
