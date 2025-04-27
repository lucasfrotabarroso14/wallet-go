package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
	"wallet-fc/internal/entity"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	ClientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255) , name  varchar(255), email  varchar(255), created_at date, updated_at date)")
	s.ClientDB = NewClientDB(db)

}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("teste", "teste@gmail.com")
	s.ClientDB.Save(client)
	clientDB, err := s.ClientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)

}

func (s *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		ID:    "1",
		Name:  "teste",
		Email: "teste@gmail.com",
	}
	err := s.ClientDB.Save(client)
	s.Nil(err)
}
