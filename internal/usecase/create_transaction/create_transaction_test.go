package create_transaction

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"wallet-fc/internal/entity"
	event2 "wallet-fc/internal/event"
	"wallet-fc/pkg/events"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (c *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := c.Called(transaction)
	return args.Error(0)
}

type ClientGatewayMock2 struct {
	mock.Mock
}

func (c *ClientGatewayMock2) Get(id string) (*entity.Client, error) {
	args := c.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}
func (c *ClientGatewayMock2) Save(client *entity.Client) error {
	args := c.Called(client)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (c *AccountGatewayMock) Save(account *entity.Account) error {
	args := c.Called(account)
	return args.Error(0)
}
func (c *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := c.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client01, _ := entity.NewClient("client01", "client01@teste.com")
	account01 := entity.NewAccount(client01)
	account01.Credit(1000)

	client02, _ := entity.NewClient("client02", "client02@teste.com")
	account02 := entity.NewAccount(client02)
	account02.Credit(1000)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindByID", account01.ID).Return(account01, nil)
	mockAccount.On("FindByID", account02.ID).Return(account02, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDTO := CreateTransactionInputDTO{
		AccountIDFrom: account01.ID,
		AccountIDTo:   account02.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	event := event2.NewTransactionCreated()

	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount, dispatcher, event)
	output, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)

}
