package create_account

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"wallet-fc/internal/entity"
)

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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("teste", "teste@gmail.com")
	clientMock := &ClientGatewayMock2{}
	clientMock.On("Get", client.ID).Return(client, nil)
	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	createAccountUseCase := NewCreateAccountUseCase(accountMock, clientMock)
	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := createAccountUseCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	//assert.Equal(t, account.ID, output.ID)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
