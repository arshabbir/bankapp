package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/arshabbir/bankapp/domain"
)

const (
	INSERT   = "INSERT"
	READ     = "READ"
	DELETE   = "DELETE"
	UPDATE   = "UPDAT"
	TRANSFER = "TRANSFER"
)

type dbClient struct {
	db *sql.DB

	tokenmap map[string]*domain.TokenMap
}

type DBCient interface {
	CreateAccount(*domain.Account) (int64, error)
	ReadAccount(AccountNumber int64) (*domain.Account, error)
	Transfer(FromAccount int64, ToAccount int64, amount int64) error
	Register(domain.User) error
	GetUser(username string, email string, password string) (string, error)
	UpdateUserToken(email string, token string) error
}

var queries map[string]*sql.Stmt

func prepareStatements(db *sql.DB) error {
	queries = make(map[string]*sql.Stmt)

	// Prepare insert statement
	insertStmt, err := db.Prepare("insert into account(owner_name, balance, currency) values ($1,$2,$3)  RETURNING id;")
	if err != nil {
		return err
	}
	// defer insertStmt.Close()
	queries[INSERT] = insertStmt

	// Prepare read statement
	readStmt, err := db.Prepare("select id, owner_name, balance, created_at, currency from  account where id=$1")
	if err != nil {
		return err
	}
	// defer insertStmt.Close()
	queries[READ] = readStmt

	// Prepare read statement
	updateStmt, err := db.Prepare("update  account set balance=balance+$1 where id=$2")
	if err != nil {
		return err
	}
	// defer insertStmt.Close()
	queries[UPDATE] = updateStmt

	return nil
}

func getStatement(idx string) *sql.Stmt {
	stmt, ok := queries[idx]
	if !ok {
		return nil
	}
	return stmt
}
func NewDBClient(host string, port int, user string, password string, dbname string) DBCient {
	// connection string
	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println("error opening  database ", psqlconn)
		return nil
	}

	// check db
	err = db.Ping()
	if err != nil {
		log.Println("error connecting to database ", err.Error())
		return nil
	}

	//	 fmt.Println("Connected!")
	log.Println("Db connection successful")
	if err := prepareStatements(db); err != nil {
		log.Println("error while preparing the statements", err.Error())
		return nil
	}
	log.Println("Successfuly prepared the statements")
	return &dbClient{db: db, tokenmap: make(map[string]*domain.TokenMap)}
}

func (c *dbClient) Register(user domain.User) error {
	// Add the user temporarely to hashmap
	_, ok := c.tokenmap[user.Email]
	if !ok {
		// Register the user
		c.tokenmap[user.Email] = &domain.TokenMap{
			U: &domain.User{Username: user.Username, Password: user.Password, Email: user.Email},
			T: &domain.LoginReponse{},
		}
		return nil
	}
	return errors.New("user already exists")

}

func (c *dbClient) GetUser(username string, email string, password string) (string, error) {
	v, ok := c.tokenmap[email]
	if !ok {
		return "", errors.New("user not exists " + email)
	}
	if v.T.Token == "" {
		return "", nil
	}
	if v.U.Username == username && v.U.Password == password {
		return v.T.Token, nil
	}
	return v.T.Token, errors.New("wrong username or password")
}

func (c *dbClient) UpdateUserToken(email string, token string) error {
	v, ok := c.tokenmap[email]
	if !ok {
		return errors.New("user not exists " + email)
	}
	v.T.Token = token
	return nil
}

func (c *dbClient) CreateAccount(acc *domain.Account) (int64, error) {
	stmt := getStatement(INSERT)
	//	defer stmt.Close()

	var id int64

	err := stmt.QueryRowContext(context.Background(), acc.Owner, acc.Balance, acc.Currency).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil

}
func (c *dbClient) ReadAccount(AccountNumber int64) (*domain.Account, error) {
	if AccountNumber < 0 {
		return nil, errors.New("invalied account number")
	}
	stmt := getStatement(READ)
	if stmt == nil {
		return nil, errors.New("error in prepare statement ")
	}
	//	defer stmt.Close()

	acc := domain.Account{}

	log.Println("Reading the account number ...... ", AccountNumber)
	rows, err := stmt.QueryContext(context.Background(), AccountNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for rows.Next() {
		err := rows.Scan(&acc.Id, &acc.Owner, &acc.Balance, &acc.Created, &acc.Currency)
		if err != nil {
			log.Println("error in row scan ......................................")
			return nil, err
		}
	}
	log.Println("REturning...............................", acc)

	return &acc, nil

}

func (c *dbClient) Transfer(FromAccount int64, ToAccount int64, amount int64) error {
	stmt := getStatement(UPDATE)
	//	defer stmt.Close()

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	// Dedut the amount from the FromAccount
	_, err = stmt.ExecContext(context.Background(), -amount, FromAccount)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Add the amount from the ToAccount
	_, err = stmt.ExecContext(context.Background(), amount, ToAccount)
	if err != nil {
		tx.Rollback()
		return err
	}
	log.Println("Commiting the Transaction")

	tx.Commit()
	return nil
}
