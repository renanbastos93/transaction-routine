// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package store

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :execresult
INSERT INTO Accounts (
  id, document_number, created_at
) VALUES (
  ?, ?, ?
)
`

type CreateAccountParams struct {
	ID             string
	DocumentNumber string
	CreatedAt      int64
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAccount, arg.ID, arg.DocumentNumber, arg.CreatedAt)
}

const createTransaction = `-- name: CreateTransaction :execresult
INSERT INTO Transactions (
  id, account_id, operation_type_id, amout, event_date
) VALUES (
  ?, ?, ?, ?, ?
)
`

type CreateTransactionParams struct {
	ID              string
	AccountID       string
	OperationTypeID string
	Amout           sql.NullFloat64
	EventDate       int64
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTransaction,
		arg.ID,
		arg.AccountID,
		arg.OperationTypeID,
		arg.Amout,
		arg.EventDate,
	)
}

const deleteAccountById = `-- name: DeleteAccountById :exec
UPDATE Accounts SET deleted_at = ? WHERE id = ?
`

type DeleteAccountByIdParams struct {
	DeletedAt int64
	ID        string
}

func (q *Queries) DeleteAccountById(ctx context.Context, arg DeleteAccountByIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteAccountById, arg.DeletedAt, arg.ID)
	return err
}

const getAllAccounts = `-- name: GetAllAccounts :many
SELECT id, document_number, created_at, deleted_at FROM Accounts
`

func (q *Queries) GetAllAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getAllAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.DocumentNumber,
			&i.CreatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllActiveOperations = `-- name: GetAllActiveOperations :many
SELECT id, description, is_active FROM OperationTypes WHERE is_active=1
`

func (q *Queries) GetAllActiveOperations(ctx context.Context) ([]Operationtype, error) {
	rows, err := q.db.QueryContext(ctx, getAllActiveOperations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Operationtype
	for rows.Next() {
		var i Operationtype
		if err := rows.Scan(&i.ID, &i.Description, &i.IsActive); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTransactions = `-- name: GetAllTransactions :many
SELECT id, account_id, operation_type_id, amout, event_date FROM Transactions
`

func (q *Queries) GetAllTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getAllTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.OperationTypeID,
			&i.Amout,
			&i.EventDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTransactionsByAccountId = `-- name: GetAllTransactionsByAccountId :many
SELECT id, account_id, operation_type_id, amout, event_date FROM Transactions WHERE account_id=?
`

func (q *Queries) GetAllTransactionsByAccountId(ctx context.Context, accountID string) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getAllTransactionsByAccountId, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.OperationTypeID,
			&i.Amout,
			&i.EventDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTransactionsByOperationTypeId = `-- name: GetAllTransactionsByOperationTypeId :many
SELECT id, account_id, operation_type_id, amout, event_date FROM Transactions WHERE operation_type_id=?
`

func (q *Queries) GetAllTransactionsByOperationTypeId(ctx context.Context, operationTypeID string) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getAllTransactionsByOperationTypeId, operationTypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.OperationTypeID,
			&i.Amout,
			&i.EventDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTransactionsByOperationTypeIdAndAccountId = `-- name: GetAllTransactionsByOperationTypeIdAndAccountId :many
SELECT id, account_id, operation_type_id, amout, event_date FROM Transactions WHERE operation_type_id=? AND account_id=?
`

type GetAllTransactionsByOperationTypeIdAndAccountIdParams struct {
	OperationTypeID string
	AccountID       string
}

func (q *Queries) GetAllTransactionsByOperationTypeIdAndAccountId(ctx context.Context, arg GetAllTransactionsByOperationTypeIdAndAccountIdParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getAllTransactionsByOperationTypeIdAndAccountId, arg.OperationTypeID, arg.AccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.OperationTypeID,
			&i.Amout,
			&i.EventDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneAccountById = `-- name: GetOneAccountById :one
SELECT id, document_number, created_at, deleted_at FROM Accounts
WHERE id = ?
`

func (q *Queries) GetOneAccountById(ctx context.Context, id string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getOneAccountById, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.DocumentNumber,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getOneOperationById = `-- name: GetOneOperationById :one
SELECT id, description, is_active FROM OperationTypes WHERE id=?
`

func (q *Queries) GetOneOperationById(ctx context.Context, id string) (Operationtype, error) {
	row := q.db.QueryRowContext(ctx, getOneOperationById, id)
	var i Operationtype
	err := row.Scan(&i.ID, &i.Description, &i.IsActive)
	return i, err
}

const getOneTransactionById = `-- name: GetOneTransactionById :one
SELECT id, account_id, operation_type_id, amout, event_date FROM Transactions WHERE id=?
`

func (q *Queries) GetOneTransactionById(ctx context.Context, id string) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getOneTransactionById, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.OperationTypeID,
		&i.Amout,
		&i.EventDate,
	)
	return i, err
}
