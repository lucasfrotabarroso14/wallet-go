package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewClient(t *testing.T) {

	client, err := NewClient("teste", "teste@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "teste", client.Name)
	assert.Equal(t, "teste@gmail.com", client.Email)

}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("teste", "teste@gmail.com")
	err := client.Update("teste_atualizado", "teste@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, "teste_atualizado", client.Name)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("teste", "teste@gmail.com")
	err := client.Update("", "")
	assert.NotNil(t, err)
	assert.Equal(t, "teste", client.Name)
	assert.Equal(t, "teste@gmail.com", client.Email)
}
