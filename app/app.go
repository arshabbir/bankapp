package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/arshabbir/bankapp/config"
	"github.com/arshabbir/bankapp/controller"
	"github.com/arshabbir/bankapp/dao"
	"github.com/arshabbir/bankapp/service"
	"github.com/gorilla/mux"
)

type bankApp struct {
	bController controller.BankController
	cfg         *config.Config
}

type BankApp interface {
	StartApp() error
}

func NewApp(cfg *config.Config) BankApp {
	return &bankApp{cfg: cfg}
}

func (a *bankApp) StartApp() error {
	var ctrl controller.BankController
	if ctrl = controller.NewBankController(service.NewBankService(dao.NewDBClient(a.cfg.Host, a.cfg.Port, a.cfg.User, a.cfg.Password, a.cfg.Dbname))); ctrl == nil {
		log.Fatal("error creating the app")
		return errors.New(("error creating the app"))
	}

	log.Println("Serving on port : ", a.cfg.AppPort)

	http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.AppPort), registerRoutes(ctrl))

	// TODO
	// Test the ping and create account

	//  Test the app creating the DB connection to pgsql
	return nil

}

func registerRoutes(ctrl controller.BankController) *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/ping", ctrl.Ping)
	mux.HandleFunc("/create", ctrl.CreateAccount)
	mux.HandleFunc("/read/{id}", ctrl.ReadAccount)
	mux.HandleFunc("/delete/{id}", ctrl.DeleteAccount)
	mux.HandleFunc("/update", ctrl.UpdateAccount)
	mux.HandleFunc("/transfer", ctrl.Transfer)

	return mux
}
