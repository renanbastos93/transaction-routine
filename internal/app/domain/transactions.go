package domain

import (
	"time"

	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
	"github.com/renanbastos93/transaction-routine/internal/app/store"
	"github.com/renanbastos93/transaction-routine/pkg/mysql"
)

// TODO in and out
type TransactionOut struct {
	weaver.AutoMarshal
	ID              string  `json:"id,omitempty"`
	AccountID       string  `json:"account_id,omitempty"`
	OperationTypeID string  `json:"operation_type_id,omitempty"`
	Amout           float64 `json:"amout,omitempty"`
	EventDate       int64   `json:"event_date,omitempty"`
}

type TransactionIn struct {
	weaver.AutoMarshal
	AccountID       string  `json:"account_id,omitempty"`
	OperationTypeID string  `json:"operation_type_id,omitempty"`
	Amout           float64 `json:"amout,omitempty"`
}

func (e TransactionOut) FromStore(in store.Transaction) TransactionOut {
	e.ID = in.ID
	e.AccountID = in.AccountID
	e.OperationTypeID = in.OperationTypeID
	e.Amout = in.Amout.Float64
	e.EventDate = in.EventDate
	return e
}

func (e TransactionIn) ToStore() (params store.CreateTransactionParams) {
	params.ID = uuid.New().String()
	params.EventDate = time.Now().UTC().UnixMilli()
	params.Amout = mysql.NewNullFloat64(e.Amout)
	params.AccountID = e.AccountID
	params.OperationTypeID = e.OperationTypeID
	return params
}

type AllTransactions []TransactionOut

func (e AllTransactions) FromStore(in []store.Transaction) AllTransactions {
	e = make(AllTransactions, 0, len(in))
	for _, v := range in {
		e = append(e, TransactionOut{}.FromStore(v))
	}
	return e
}
