package mocks

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/renanbastos93/transaction-routine/internal/app/domain"
)

type appFake struct {
	ErrCreate           error
	ErrTrasaction       error
	ErrGetUserById      error
	ErrGetAllUsers      error
	ErrGetAllOperations error
	ErrThereIsOperation error
}

func (e appFake) CreateUser(ctx context.Context, in domain.AccountIn) (err error) {
	return e.ErrCreate
}

func (e appFake) GetUserById(ctx context.Context, id string) (out domain.AccountOut, err error) {
	return out, e.ErrGetUserById
}

func (e appFake) GetAllUsers(ctx context.Context) (out domain.AllAccounts, err error) {
	return out, e.ErrGetAllUsers
}

func (e appFake) SaveTransaction(ctx context.Context, in domain.TransactionIn) (ok bool, err error) {
	return ok, e.ErrTrasaction
}

func (e appFake) GetAllOperations(ctx context.Context) (out domain.AllOperations, err error) {
	return out, e.ErrGetAllOperations
}

func (e appFake) ThereIsOperationTypById(ctx context.Context, id string) (err error) {
	return e.ErrThereIsOperation
}

func NewMockApp() (a appFake) {
	return a
}

func NewMockCreateUserError() (a appFake) {
	a.ErrCreate = fiber.NewError(fiber.StatusInternalServerError, "anything")
	return a
}

func NewMockGetUserByIdErrorForbidden() (a appFake) {
	a.ErrGetUserById = fiber.ErrForbidden
	return a
}

func NewMockGetUserByIdErrorNotFound() (a appFake) {
	a.ErrGetUserById = fiber.NewError(fiber.StatusNotFound, "anything")
	return a
}

func NewMockGetAllAccountsWithError() (a appFake) {
	a.ErrGetAllUsers = fiber.NewError(fiber.StatusNotFound, "anything")
	return a
}

func NewMockGetAllOperationsWithError() (a appFake) {
	a.ErrGetAllOperations = fiber.NewError(fiber.StatusNotFound, "anything")
	return a
}

func NewMockThereIsOperationWithError() (a appFake) {
	a.ErrThereIsOperation = fiber.NewError(fiber.StatusBadRequest, "anything")
	return a
}

func NewMockSaveTransactionError() (a appFake) {
	a.ErrTrasaction = fiber.NewError(fiber.StatusInternalServerError, "anything")
	return a
}
