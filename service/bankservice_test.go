package service

import (
	"log"
	"testing"
	"time"

	"github.com/arshabbir/bankapp/domain"
	"github.com/stretchr/testify/suite"
)

type fakeDAO struct {
}

type DbTestSuite struct {
	suite.Suite
}

/*
CreateAccount(*domain.Account) (int64, error)
	ReadAccount(AccountNumber int64) (*domain.Account, error)
	Transfer(FromAccount int64, ToAccount int64, amount int64) error

*/

// NEED TO WORK ON TABLE DRIVEN TESTS

var fDb fakeDAO
var s BankService

func (d *DbTestSuite) SetupTest() {
	log.Println("setup executed.")
	fDb = fakeDAO{}
	s = NewBankService(&fDb)

}

func (d *DbTestSuite) TearDownTest() {
	log.Println("tear down called")
}

func TestDbTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}
func (f *fakeDAO) CreateAccount(acc *domain.Account) (int64, error) {
	return acc.Id, nil

}

func (f *fakeDAO) ReadAccount(AccountNumber int64) (*domain.Account, error) {
	return &domain.Account{Id: AccountNumber, Owner: "arshabbir", Balance: 100, Created: time.Now()}, nil

}

func (f *fakeDAO) Transfer(FromAccount int64, ToAccount int64, amount int64) error {
	return nil

}

func (d *DbTestSuite) TestCreateAccountSuccess() {

	acc := &domain.Account{Id: 7, Owner: "arshabbir", Balance: 100, Currency: "USD", Created: time.Now()}

	resp, err := s.CreateAccount(acc)

	d.Nil(err)
	d.Equal(resp, acc.Id)

}

func (d *DbTestSuite) TestCreateAccountFailed() {

	resp, err := s.CreateAccount(nil)

	d.NotNil(err)
	d.Equal(resp, int64(-1))
	d.Equal(err.Error(), "Nil account")

}

func (d *DbTestSuite) TestReadAccountSuccess() {

	var accNumber int64 = 7

	resp, err := s.ReadAccount(int64(accNumber))

	d.Nil(err)
	d.Equal(accNumber, resp.Id)
}

func (d *DbTestSuite) TestReadAccountFailed() {

	accNumber := -1

	resp, err := s.ReadAccount(int64(accNumber))

	d.NotNil(err)
	d.Nil(resp)
	d.Equal(err.Error(), "account number cannot be negitive")
}
