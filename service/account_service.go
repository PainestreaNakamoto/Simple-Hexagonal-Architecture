package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func InitializeAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountRepsonse, error) {
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5000")
	}
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("Account Type should be saving or checking")
	}
	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}
	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	response := AccountRepsonse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountRepsonse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	response := []AccountRepsonse{}
	for _, item := range accounts {
		response = append(response, AccountRepsonse{
			AccountID:   item.AccountID,
			OpeningDate: item.OpeningDate,
			AccountType: item.AccountType,
			Amount:      item.Amount,
			Status:      item.Status,
		})
	}

	return response, nil
}
