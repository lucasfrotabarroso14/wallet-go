package create_client

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"wallet-fc/internal/entity"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (c *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := c.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}
func (c *ClientGatewayMock) Save(client *entity.Client) error {
	args := c.Called(client)
	return args.Error(0)
}

func TestNewCreateClientUseCase(t *testing.T) {
	inputCreateClient := CreateClientInputDTO{
		Name:  "teste",
		Email: "teste@teste.com",
	}
	clientGatewayMock := new(ClientGatewayMock)
	clientGatewayMock.On("Save", mock.Anything).Return(nil)
	createClientUseCase := NewCreateClientUseCase(clientGatewayMock)
	output, err := createClientUseCase.Execute(inputCreateClient)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, inputCreateClient.Name, output.Name)
	assert.Equal(t, inputCreateClient.Email, output.Email)
	clientGatewayMock.AssertExpectations(t)

}
