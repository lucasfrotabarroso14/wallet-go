package database

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
	"wallet-fc/internal/entity"
)

type TransactionDBTestSuite struct {
	DB            *sql.DB
	client        *entity.Client
	client02      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	TransactionDB *TransactionDB
	suite.Suite
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.DB = db
	db.Exec("CREATE TABLE clients (id varchar(255) , name  varchar(255), email  varchar(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255) , client_id  varchar(255), balance  float, created_at date, updated_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255) , account_id_from  varchar(255), account_id_to  varchar(255), amount  float, created_at date, updated_at date)")
	client, err := entity.NewClient("john", "j@j.com")
	s.Nil(err)
	s.client = client

	client02, err := entity.NewClient("john2", "j@j.com")
	s.Nil(err)
	s.client02 = client02

	//creating account
	accountFrom := entity.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.client02)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.TransactionDB = NewTransactionDB(db)

}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.DB.Close()
	s.DB.Exec("DROP TABLE clients")
	s.DB.Exec("DROP TABLE accounts")
	s.DB.Exec("DROP TABLE transactions")
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.TransactionDB.Create(transaction)
	s.Nil(err)
}
