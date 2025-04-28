package create_transaction

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"wallet-fc/internal/entity"
	event2 "wallet-fc/internal/event"
	"wallet-fc/internal/usecase/mocks"
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

	mockTransaction := &TransactionGatewayMock{}

	uowMock := &mocks.UowMock{}
	uowMock.On("Do", mock.Anything, mock.Anything).Return(nil)
	inputDTO := CreateTransactionInputDTO{
		AccountIDFrom: account01.ID,
		AccountIDTo:   account02.ID,
		Amount:        100,
	}
	dispatcher := events.NewEventDispatcher()
	event := event2.NewTransactionCreated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(uowMock, dispatcher, event)
	err := uc.Execute(ctx, inputDTO)
	assert.Nil(t, err)
	uowMock.AssertExpectations(t)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	uowMock.AssertNumberOfCalls(t, "Do", 1)

}
