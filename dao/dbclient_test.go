package dao

import (
	"log"
	"testing"
	"time"

	"github.com/arshabbir/bankapp/domain"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	log.Println("Running TestCreateaccount")
	db := NewDBClient("localhost", 5432, "postgres", "password", "bank")
	// NEED TO FIX THE EMPTY AND NOT NULL entry test
	d := domain.Account{Owner: "", Currency: "USD", Created: time.Now()}
	id, err := db.CreateAccount(&d)
	require.NotEmpty(t, id)
	require.Nil(t, err)
}

func TestReadAccount(t *testing.T) {
	log.Println("Running TestReadAccount")
	db := NewDBClient("localhost", 5432, "postgres", "password", "bank")
	// NEED TO FIX THE EMPTY AND NOT NULL entry test
	d := domain.Account{Owner: "test", Currency: "USD", Created: time.Now(), Balance: 200}
	id, err := db.CreateAccount(&d)
	if err != nil {
		log.Fatal("error while setup")
	}

	res, err := db.ReadAccount(id)
	require.Nil(t, err)
	require.Equal(t, res.Balance, d.Balance)
	require.Equal(t, id, res.Id)
	require.Equal(t, d.Currency, res.Currency)

	res, err = db.ReadAccount(-1)
	require.NotNil(t, err)
	require.Nil(t, res)

}

func TestUpdateAccount(t *testing.T) {

}

func TestDeleteAccount(t *testing.T) {

}

func TestTransfer(t *testing.T) {

	// Setup
	db := NewDBClient("localhost", 5432, "postgres", "password", "bank")
	// NEED TO FIX THE EMPTY AND NOT NULL entry test
	d1 := domain.Account{Owner: "test1", Currency: "USD", Created: time.Now(), Balance: 1500}
	id1, err := db.CreateAccount(&d1)
	if err != nil {
		log.Fatal("error while setup")
	}

	// NEED TO FIX THE EMPTY AND NOT NULL entry test
	d2 := domain.Account{Owner: "test2", Currency: "USD", Created: time.Now(), Balance: 1000}
	id2, err := db.CreateAccount(&d2)
	if err != nil {
		log.Fatal("error while setup")
	}
	log.Println("Running TestTransfer")

	var amount int64 = 100

	errChan := make(chan error)
	// INitiate transfer
	go func(chan error) {
		for i := 0; i < 5; i++ {
			err = db.Transfer(id1, id2, amount)
			require.Nil(t, err)
		}
		errChan <- err
	}(errChan)

	v := <-errChan
	require.Nil(t, v)

	acc1, err1 := db.ReadAccount(id1)
	acc2, err2 := db.ReadAccount(id2)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Equal(t, acc2.Balance, d2.Balance+5*amount)
	require.Equal(t, acc1.Balance, d1.Balance-5*amount)

}
