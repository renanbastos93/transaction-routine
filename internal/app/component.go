package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ServiceWeaver/weaver"
	"github.com/renanbastos93/transaction-routine/internal/app/domain"
	"github.com/renanbastos93/transaction-routine/internal/app/store"
	"github.com/renanbastos93/transaction-routine/pkg/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type Component interface {
	CreateUser(ctx context.Context, in domain.AccountIn) (err error)
	GetUserById(ctx context.Context, id string) (out domain.AccountOut, err error)
	GetAllUsers(ctx context.Context) (out domain.AllAccounts, err error)
	SaveTransaction(ctx context.Context, in domain.TransactionIn) (ok bool, err error)
	GetAllOperations(ctx context.Context) (out domain.AllOperations, err error)
	ThereIsOperationTypById(ctx context.Context, id string) (err error)
}

type Config struct {
	Driver string
	Source string
}

type implapp struct {
	weaver.Implements[Component]
	weaver.WithConfig[Config]
	db *store.Queries
}

func (e *implapp) Init(ctx context.Context) error {
	db, err := sql.Open(e.Config().Driver, e.Config().Source)
	if err != nil {
		return fmt.Errorf("not open: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping: %w", err)
	}

	e.db = store.New(db)
	return nil
}

func (e implapp) CreateUser(ctx context.Context, in domain.AccountIn) (err error) {
	account := in.ToStore()
	_, err = e.db.CreateAccount(ctx, account)
	if err != nil {
		return err
	}
	return nil
}

func (e implapp) GetUserById(ctx context.Context, id string) (out domain.AccountOut, err error) {
	account, err := e.db.GetOneAccountById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return out, ErrNotFoundAccount
		}
		return out, err
	}

	return out.FromStore(account), nil
}

func (e implapp) GetAllUsers(ctx context.Context) (out domain.AllAccounts, err error) {
	accounts, err := e.db.GetAllAccounts(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(accounts), nil
}

func (e implapp) GetAllOperations(ctx context.Context) (out domain.AllOperations, err error) {
	operations, err := e.db.GetAllActiveOperations(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(operations), nil
}

func (e implapp) SaveTransaction(ctx context.Context, in domain.TransactionIn) (ok bool, err error) {
	operationType, err := e.db.GetOneOperationById(ctx, in.OperationTypeID)
	if err != nil {
		return ok, ErrNotFoundOperations
	}

	if !operationType.IsActive {
		return ok, ErrInactiveOperation
	}

	storageData := in.ToStore()

	switch {
	case strings.Contains(operationType.Description, domain.OperationCash) || strings.Contains(operationType.Description, domain.OperationWithdrawal):
		if in.Amout > 0 {
			storageData.Amout = mysql.NewNullFloat64(in.Amout * -1)
			storageData.Balance = storageData.Amout.Float64
		}

	case strings.Contains(operationType.Description, domain.OperationPay):
		if in.Amout < 0 {
			storageData.Amout = mysql.NewNullFloat64(in.Amout * -1)
		}

		transactions, err := e.db.GetAllTrasactionsNotPaymentByAccountIdAndLessZeroAndEventDateFilter(ctx, in.AccountID)
		if err != nil {
			return false, fmt.Errorf("cannot filter the last balance")
		}

		if len(transactions) == 0 {
			storageData.Balance = storageData.Amout.Float64
		} else {
			// TODO: create another func
			var updateTransactions = make([]store.UpdateTransactionByIdParams, 0, len(transactions))
			for _, v := range transactions {
				v.Balance = v.Balance + (v.Balance * -1)
				storageData.Balance = v.Balance

				params := store.UpdateTransactionByIdParams{
					ID:      v.ID,
					Balance: v.Balance,
				}

				// TODO: we need to improve that
				if v.Balance > 0 {
					params.Balance = 0
				}

				updateTransactions = append(updateTransactions, params)
			}

			// TODO: create another func
			go func() {
				for _, v := range updateTransactions {
					err := e.db.UpdateTransactionById(ctx, v)
					if err != nil {
						fmt.Println("err to update transaction after calculate balance", err.Error())
					}
				}
			}()
		}
	}

	_, err = e.db.CreateTransaction(ctx, storageData)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e implapp) ThereIsOperationTypById(ctx context.Context, id string) (err error) {
	operation, err := e.db.GetOneOperationById(ctx, id)
	if err != nil {
		return ErrNotFoundOperations
	}

	if !operation.IsActive {
		return ErrNotFoundOperations
	}

	return nil
}
