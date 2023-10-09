-- name: GetAllAccounts :many
SELECT * FROM Accounts;

-- name: GetOneAccountById :one
SELECT * FROM Accounts
WHERE id = ?;

-- name: CreateAccount :execresult
INSERT INTO Accounts (
  id, document_number, created_at
) VALUES (
  ?, ?, ?
);

-- name: DeleteAccountById :exec
UPDATE Accounts SET deleted_at = ? WHERE id = ?;

-- name: GetAllActiveOperations :many
SELECT * FROM OperationTypes WHERE is_active=1;

-- name: GetOneOperationById :one
SELECT * FROM OperationTypes WHERE id=?;

-- name: CreateTransaction :execresult
INSERT INTO Transactions (
  id, account_id, operation_type_id, amout, event_date
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: GetAllTransactions :many
SELECT * FROM Transactions;

-- name: GetOneTransactionById :one
SELECT * FROM Transactions WHERE id=?;

-- name: GetAllTransactionsByAccountId :many
SELECT * FROM Transactions WHERE account_id=?;

-- name: GetAllTransactionsByOperationTypeId :many
SELECT * FROM Transactions WHERE operation_type_id=?;

-- name: GetAllTransactionsByOperationTypeIdAndAccountId :many
SELECT * FROM Transactions WHERE operation_type_id=? AND account_id=?;
