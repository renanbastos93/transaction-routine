package bff

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/renanbastos93/transaction-routine/internal/app/domain"
)

func (e implBFF) RegisterAccounts(router fiber.Router) {
	router.Post("/accounts", e.createAccount)
	router.Get("/accounts/:id", e.getAccountById)
	router.Post("/transactions", e.saveTransaction)
}

func (e implBFF) createAccount(c *fiber.Ctx) (err error) {
	var body domain.AccountIn

	err = c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "body invalid")
	}

	err = e.app.Get().CreateUser(context.Background(), body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (e implBFF) getAccountById(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	if id == "" {
		return fiber.ErrBadRequest
	}

	out, err := e.app.Get().GetUserById(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) saveTransaction(c *fiber.Ctx) (err error) {
	var body domain.TransactionIn

	err = c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "body invalid")
	}

	if _, err := e.app.Get().GetUserById(c.Context(), body.AccountID); err != nil {
		return fiber.ErrForbidden
	}

	if err = e.app.Get().ThereIsOperationTypById(c.Context(), body.OperationTypeID); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err = e.app.Get().SaveTransaction(context.Background(), body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}
