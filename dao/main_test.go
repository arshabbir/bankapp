package dao

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Test main executed")
	os.Exit(m.Run())

}
