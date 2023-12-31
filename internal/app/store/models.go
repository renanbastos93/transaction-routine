// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package store

import (
	"database/sql"
)

type Account struct {
	ID             string
	DocumentNumber string
	CreatedAt      int64
	DeletedAt      int64
}

type Operationtype struct {
	ID          string
	Description string
	IsActive    bool
}

type Transaction struct {
	ID              string
	AccountID       string
	OperationTypeID string
	Amout           sql.NullFloat64
	EventDate       int64
	Balance         float64
}
