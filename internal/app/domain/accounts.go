package domain

import (
	"time"

	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
	"github.com/renanbastos93/transaction-routine/internal/app/store"
)

type AccountOut struct {
	weaver.AutoMarshal
	ID             string `json:"id"`
	DocumentNumber string `json:"document_number"`
	CreatedAt      int64  `json:"created_at"`
	DeletedAt      int64  `json:"deleted_at"`
}

type AccountIn struct {
	weaver.AutoMarshal
	DocumentNumber string `json:"document_number"`
}

func (e AccountOut) FromStore(in store.Account) AccountOut {
	e.ID = in.ID
	e.DocumentNumber = in.DocumentNumber
	e.CreatedAt = in.CreatedAt
	return e
}

func (e AccountIn) ToStore() (params store.CreateAccountParams) {
	params.ID = uuid.New().String()
	params.CreatedAt = time.Now().UTC().UnixMilli()
	params.DocumentNumber = e.DocumentNumber
	return params
}

type AllAccounts []AccountOut

func (e AllAccounts) FromStore(in []store.Account) AllAccounts {
	e = make(AllAccounts, 0, len(in))
	for _, v := range in {
		e = append(e, AccountOut{}.FromStore(v))
	}
	return e
}
