package create_client

import (
	"time"
	"wallet-fc/internal/gateway"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID string
	Name string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase interface {
	ClientGateway ClientGate
}
