package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	accountFrom := &Account{ID: "123", Balance: 1000}
	accountTo := &Account{ID: "456", Balance: 500}
	amount := 200.0
	transaction, err := NewTransaction(accountFrom, accountTo, amount)
	assert.NoError(t, err)
	assert.NotNil(t, transaction)

}

func TestTransaction_Commit(t *testing.T) {
	accountFrom := &Account{ID: "123", Balance: 1000.0}
	accountTo := &Account{ID: "456", Balance: 500.0}
	amount := 200.0
	transaction, _ := NewTransaction(accountFrom, accountTo, amount)

	assert.Equal(t, 800.0, transaction.AccountFrom.Balance)
	assert.Equal(t, 700.0, transaction.AccountTo.Balance)
}
