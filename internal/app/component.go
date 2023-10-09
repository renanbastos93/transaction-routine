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
	// if compra && saque { save negative number } else {save positive}
	operationType, err := e.db.GetOneOperationById(ctx, in.OperationTypeID)
	if err != nil {
		return ok, ErrNotFoundOperations
	}

	if !operationType.IsActive {
		return ok, ErrInactiveOperation
	}

	switch {
	case strings.Contains(operationType.Description, domain.OperationCash) || strings.Contains(operationType.Description, domain.OperationWithdrawal):
		if in.Amout > 0 {
			in.Amout *= -1
		}
	case strings.Contains(operationType.Description, domain.OperationPay):
		if in.Amout < 0 {
			in.Amout *= -1
		}
	}

	_, err = e.db.CreateTransaction(ctx, in.ToStore())
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
