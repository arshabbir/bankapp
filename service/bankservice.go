package service

import (
	"errors"
	"time"

	"github.com/arshabbir/bankapp/config"
	"github.com/arshabbir/bankapp/dao"
	"github.com/arshabbir/bankapp/domain"
	"github.com/dgrijalva/jwt-go"
)

type bankService struct {
	dbClient dao.DBCient
}

type BankService interface {
	CreateAccount(*domain.Account) (int64, error)
	ReadAccount(AccountNumber int64) (*domain.Account, error)
	UpdateAccount(*domain.Account) error
	DeleteAccount(AccountNumber int64) error
	Transfer(FromAccountID int64, ToAccountID int64, amount int64) error
	Register(domain.User) error
	Login(domain.TokenRequest) (*domain.TokenReponse, error)
}

func NewBankService(dbClient dao.DBCient) BankService {
	return &bankService{dbClient: dbClient}
}

func (c *bankService) Register(user domain.User) error {
	return c.dbClient.Register(user)
}

func (c *bankService) Login(r domain.TokenRequest) (*domain.TokenReponse, error) {

	token, err := c.dbClient.CheckUser(r.Username, r.Email, r.Password)
	if err != nil {
		return nil, err
	}

	// check if token exists
	if token != "" {
		return &domain.TokenReponse{Token: token, Username: r.Username}, nil
	}

	// Generate the token and respond
	token, err = generateJWT(r.Username, r.Email)
	if err != nil {
		return nil, err
	}

	return &domain.TokenReponse{Username: r.Username, Token: token}, nil
}

func (c *bankService) CreateAccount(acc *domain.Account) (int64, error) {
	if acc == nil {
		return -1, errors.New("Nil account")
	}
	return c.dbClient.CreateAccount(acc)
}
func (c *bankService) ReadAccount(AccountNumber int64) (*domain.Account, error) {
	if AccountNumber < 0 {
		return nil, errors.New("account number cannot be negitive")
	}
	return c.dbClient.ReadAccount(AccountNumber)

}
func (c *bankService) UpdateAccount(acc *domain.Account) error {
	return nil
}

func (c *bankService) DeleteAccount(AccountNumber int64) error {
	return nil
}

func (c *bankService) Transfer(FromAccount int64, ToAccount int64, amount int64) error {
	return c.dbClient.Transfer(FromAccount, ToAccount, amount)
}

func generateJWT(username string, email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &domain.JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(config.GlobalConf.JWTKEY)
	return
}
func validateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&domain.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GlobalConf.JWTKEY), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*domain.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
