package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/arshabbir/bankapp/config"
	"github.com/arshabbir/bankapp/domain"
	"github.com/arshabbir/bankapp/service"
	"github.com/arshabbir/bankapp/util"
	"github.com/gorilla/mux"
)

type controller struct {
	bankService service.BankService
}
type BankController interface {
	CreateAccount(http.ResponseWriter, *http.Request)
	ReadAccount(http.ResponseWriter, *http.Request)
	UpdateAccount(http.ResponseWriter, *http.Request)
	DeleteAccount(http.ResponseWriter, *http.Request)
	Transfer(http.ResponseWriter, *http.Request)
	Start() error
	Ping(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

func NewBankController(bService service.BankService) BankController {
	return &controller{bankService: bService}
}

func (c *controller) Login(w http.ResponseWriter, r *http.Request) {

	// TODO : Get username passwrod from token request & pass to Login service function to get the jwt token

	l := domain.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusBadRequest, Message: "Invalid request " + err.Error()}); err != nil {
			log.Fatal("error while sending the response")
		}
		return
	}
	defer r.Body.Close()
	resp, err := c.bankService.Login(l)
	if err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusInternalServerError, Message: "unable to login " + err.Error()}); err != nil {
			log.Fatal("error while sending the response")
		}
		return
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatal("error while sending the response ", err)
	}

}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	u := domain.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusBadRequest, Message: "Invalid request " + err.Error()}); err != nil {
			log.Fatal("error while sending the response")
		}
		return
	}

	if err := c.bankService.Register(u); err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusInternalServerError, Message: "unable to register " + err.Error()}); err != nil {
			log.Fatal("error while sending the response")
		}
		return
	}

	if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusOK, Message: fmt.Sprintf("Registration successful : %s", u.Username)}); err != nil {
		log.Fatal("error while sending the response ", err)
	}

}

func (c *controller) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (c *controller) Start() error {
	log.Println("Bank Controller serving on port : ", config.GlobalConf.AppPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.GlobalConf.AppPort), registerRoutes(c))
}
func (c *controller) CreateAccount(w http.ResponseWriter, r *http.Request) {
	acc := domain.Account{}
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusBadRequest, Message: "Invalid request " + err.Error()}); err != nil {
			log.Fatal("error while sending the response")
		}
		return
	}
	resp, err := c.bankService.CreateAccount(&acc)
	if err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusInternalServerError, Message: "Internal server error " + err.Error()}); err != nil {
			log.Fatal("error while sending the response ", err)
		}
		return
	}
	if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusOK, Message: fmt.Sprintf("Account creation successful : %d", resp)}); err != nil {
		log.Fatal("error while sending the response ", err)
	}
}
func (c *controller) ReadAccount(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	nID, err := strconv.Atoi(id)
	if err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusBadRequest, Message: "Account Number is mandatory "}); err != nil {
			log.Fatal("error while sending the response ", err)
		}
		return

	}
	acc, err := c.bankService.ReadAccount(int64(nID))
	if err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusInternalServerError, Message: "Internal server error " + err.Error()}); err != nil {
			log.Fatal("error while sending the response ", err)
		}
		return
	}

	if err = json.NewEncoder(w).Encode(&acc); err != nil {
		log.Fatal("error while sending the response")
	}

}
func (c *controller) UpdateAccount(w http.ResponseWriter, r *http.Request) {

}

func (c *controller) DeleteAccount(w http.ResponseWriter, r *http.Request) {
}

func (c *controller) Transfer(w http.ResponseWriter, r *http.Request) {
	tx := domain.Transactions{}
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusBadRequest, Message: "Invalid request " + err.Error()}); err != nil {
			log.Fatal("error while sending the response")
		}
		return
	}
	err := c.bankService.Transfer(tx.FromAccount, tx.ToAccount, tx.Amount)
	if err != nil {
		if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusInternalServerError, Message: "Transaction FAILED :  " + err.Error()}); err != nil {
			log.Fatal("error while sending the response")
		}
		return
	}
	if err := json.NewEncoder(w).Encode(&util.ApiError{Statuscode: http.StatusOK, Message: "Transaction successful "}); err != nil {
		log.Fatal("error while sending the response ", err)
	}

}

func registerRoutes(ctrl BankController) *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/ping", ctrl.Ping)
	mux.HandleFunc("/create", ctrl.CreateAccount)
	mux.HandleFunc("/read/{id}", ctrl.ReadAccount)
	mux.HandleFunc("/delete/{id}", ctrl.DeleteAccount)
	mux.HandleFunc("/update", ctrl.UpdateAccount)
	mux.HandleFunc("/transfer", ctrl.Transfer)
	mux.HandleFunc("/login", ctrl.Login)
	mux.HandleFunc("/register", ctrl.Register)

	return mux
}
