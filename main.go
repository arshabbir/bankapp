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
	conf := config.Config{}
	c, err := config.ReadConfig(".", &conf)

	if err != nil {
		log.Fatal("error while reading the config file")
	}
	config.GlobalConf = c
	log.Println("Configuration Loaded ", c)
	db := dao.NewDBClient(c.Host, c.Port, c.User, c.Password, c.Dbname)
	if db == nil {
		log.Fatal("error connting to db")
	}
	bapp := app.NewApp(controller.NewBankController(service.NewBankService(db)))
	log.Println("Starting Banking application..")
	bapp.StartApp()
}
