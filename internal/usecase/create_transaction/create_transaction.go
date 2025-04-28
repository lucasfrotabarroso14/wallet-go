package create_transaction

import (
	"context"
	"wallet-fc/internal/entity"
	"wallet-fc/internal/gateway"
	"wallet-fc/pkg/events"
	"wallet-fc/pkg/uow"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	eventDispatcher    events.EventDispatcherInterface
	transactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(Uow uow.UowInterface, eventDispatcher events.EventDispatcherInterface, transactionCreated events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                Uow,
		eventDispatcher:    eventDispatcher,
		transactionCreated: transactionCreated,
	}
}
func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDTO) error {
	output := &CreateTransactionOutputDTO{}
	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := uc.getAccountRepository(ctx)
		transactionRepository := uc.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountIDFrom)
		if err != nil {
			return err
		}
		accountTo, err := accountRepository.FindByID(input.AccountIDTo)
		if err != nil {
			return err
		}
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}
		output.ID = transaction.ID
		output.AccountIDFrom = accountFrom.ID
		output.AccountIDTo = accountTo.ID
		output.Amount = transaction.Amount
		return nil

	})
	if err != nil {
		return err
	}

	uc.transactionCreated.SetPayload(output)
	_ = uc.eventDispatcher.Dispatch(uc.transactionCreated)

	return nil

}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
