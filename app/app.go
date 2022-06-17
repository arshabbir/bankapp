package app

import (
	"errors"
	"log"

	"github.com/arshabbir/bankapp/config"
	"github.com/arshabbir/bankapp/controller"
)

type bankApp struct {
	bController controller.BankController
	cfg         *config.Config
}

type BankApp interface {
	StartApp() error
}

func NewApp(ctrl controller.BankController, cfg *config.Config) BankApp {
	return &bankApp{cfg: cfg, bController: ctrl}
}

func (a *bankApp) StartApp() error {
	if a.bController == nil {
		log.Fatal("error creating the app")
		return errors.New(("error creating the app"))
	}

	return a.bController.Start()

}
