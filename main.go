package main

import (
	"log"

	"github.com/arshabbir/bankapp/app"
	"github.com/arshabbir/bankapp/config"
)

func main() {
	c := &config.Config{
		Dbname:   "bank",
		User:     "postgres",
		Host:     "postgres",
		Port:     5432,
		Password: "password",
		AppPort:  8080,
	}
	bapp := app.NewApp(c)
	log.Println("Starting Banking application..")
	bapp.StartApp()
}
