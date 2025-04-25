package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	client, err := NewClient("teste", "teste@gmail.com")
	assert.Nil(t, err)

	account := NewAccount(client)
	assert.NotNil(t, account)
}

func TestNewAccountWithInvalidClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestAccountCredit(t *testing.T) {
	client, _ := NewClient("teste", "teste@gmail.com")
	account := NewAccount(client)

	account.Credit(100.0)
	assert.Equal(t, 100.0, account.Balance)

}

func TestAccountDebit(t *testing.T) {
	client, _ := NewClient("teste", "teste@gmail.com")
	account := NewAccount(client)
	account.Credit(100.0)
	account.Debit(50.0)
	assert.Equal(t, 50.0, account.Balance)
}

func TestClient_AddAccount(t *testing.T) {
	client, _ := NewClient("teste", "teste@gmail.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
}
