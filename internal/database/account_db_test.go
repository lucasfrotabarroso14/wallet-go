package database

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
	"wallet-fc/internal/entity"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	AccountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255) , name  varchar(255), email  varchar(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255) , client_id  varchar(255), balance  float, created_at date, updated_at date)")
	s.AccountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("teste", "teste@gmail.com")

}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		s.client.ID,
		s.client.Name,
		s.client.Email,
		s.client.CreatedAt,
		s.client.UpdatedAt,
	)
	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
	accountDB, err := s.AccountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Balance, accountDB.Balance)
}

func TestAccountTestsSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}
