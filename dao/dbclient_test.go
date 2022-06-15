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

}

func TestUpdateAccount(t *testing.T) {

}

func TestDeleteAccount(t *testing.T) {

}

func TestTransfer(t *testing.T) {

}
