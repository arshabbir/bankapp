package service

import (
	"github.com/arshabbir/bankapp/dao"
	"github.com/arshabbir/bankapp/domain"
)

type bankService struct {
	dbClient dao.DBCient
}

type BankService interface {
	CreateAccount(*domain.Account) (string, error)
	ReadAccount(AccountNumber int64) (*domain.Account, error)
	UpdateAccount(*domain.Account) error
	DeleteAccount(AccountNumber int64) error
	Transfer(FromAccountID int64, ToAccountID int64, amount int64) error
}

func NewBankService(dbClient dao.DBCient) BankService {
	return &bankService{dbClient: dbClient}
}
func (c *bankService) CreateAccount(acc *domain.Account) (string, error) {
	return c.dbClient.CreateAccount(acc)
}
func (c *bankService) ReadAccount(AccountNumber int64) (*domain.Account, error) {
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
