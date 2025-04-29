package handler

import (
	"fmt"
	"sync"
	"wallet-fc/pkg/events"
	"wallet-fc/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := h.kafka.Publish(message, nil, "transactions"); err != nil {
		fmt.Println("Erro ao publica mensagem no kafka ->:", err)
		return
	}
	fmt.Println("TransactionCreatedKafkaHandler:", message.GetPayload())

}
