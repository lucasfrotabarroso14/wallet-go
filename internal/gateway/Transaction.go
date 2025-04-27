package gateway

import "wallet-fc/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
