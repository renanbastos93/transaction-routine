package domain

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/renanbastos93/transaction-routine/internal/app/store"
)

type Operation struct {
	weaver.AutoMarshal
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	IsActive    bool   `json:"is_active,omitempty"`
}

func (e Operation) FromStore(in store.Operationtype) Operation {
	e.ID = in.ID
	e.Description = in.Description
	e.IsActive = in.IsActive
	return e
}

type AllOperations []Operation

func (e AllOperations) FromStore(in []store.Operationtype) AllOperations {
	e = make(AllOperations, 0, len(in))
	for _, v := range in {
		e = append(e, Operation{}.FromStore(v))
	}
	return e
}

const (
	OperationCash       = "COMPRA"
	OperationWithdrawal = "SAQUE"
	OperationPay        = "PAGAMENTO"
)
