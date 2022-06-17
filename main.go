package main

import (
	"log"

	"github.com/arshabbir/bankapp/app"
	"github.com/arshabbir/bankapp/config"
	"github.com/arshabbir/bankapp/controller"
	"github.com/arshabbir/bankapp/dao"
	"github.com/arshabbir/bankapp/service"
)

func main() {
	c := &config.Config{
		Dbname:   "bank",
		User:     "postgres",
		Host:     "localhost",
		Port:     5432,
		Password: "password",
		AppPort:  8080,
	}
	bapp := app.NewApp(controller.NewBankController(service.NewBankService(dao.NewDBClient(c.Host, c.Port, c.User, c.Password, c.Dbname)), c), c)
	log.Println("Starting Banking application..")
	bapp.StartApp()
}
