package service

import (
	"errors"

	"github.com/arshabbir/bankapp/dao"
	"github.com/arshabbir/bankapp/domain"
)

type bankService struct {
	dbClient dao.DBCient
}

type BankService interface {
	CreateAccount(*domain.Account) (int64, error)
	ReadAccount(AccountNumber int64) (*domain.Account, error)
	UpdateAccount(*domain.Account) error
	DeleteAccount(AccountNumber int64) error
	Transfer(FromAccountID int64, ToAccountID int64, amount int64) error
}

func NewBankService(dbClient dao.DBCient) BankService {
	return &bankService{dbClient: dbClient}
}
func (c *bankService) CreateAccount(acc *domain.Account) (int64, error) {
	if acc == nil {
		return -1, errors.New("Nil account")
	}
	return c.dbClient.CreateAccount(acc)
}
func (c *bankService) ReadAccount(AccountNumber int64) (*domain.Account, error) {
	if AccountNumber < 0 {
		return nil, errors.New("account number cannot be negitive")
	}
	return c.dbClient.ReadAccount(AccountNumber)

}
func (c *bankService) UpdateAccount(acc *domain.Account) error {
	return nil
}

func (c *bankService) DeleteAccount(AccountNumber int64) error {
	return nil
}

func (c *bankService) Transfer(FromAccount int64, ToAccount int64, amount int64) error {
	return c.dbClient.Transfer(FromAccount, ToAccount, amount)
}
