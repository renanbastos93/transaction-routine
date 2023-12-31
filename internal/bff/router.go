package bff

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/renanbastos93/transaction-routine/internal/app/domain"
)

func (e implBFF) RegisterAccounts(router fiber.Router) {
	router.Get("/accounts", e.getAllAccounts)
	router.Post("/accounts", e.createAccount)
	router.Get("/accounts/:id", e.getAccountById)
	router.Get("/operations", e.getAllOperations)
	router.Post("/transactions", e.saveTransaction)
}

func (e implBFF) getAllAccounts(c *fiber.Ctx) (err error) {
	out, err := e.app.Get().GetAllUsers(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) getAllOperations(c *fiber.Ctx) (err error) {
	out, err := e.app.Get().GetAllOperations(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(out)
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
	out, err := e.app.Get().GetUserById(c.Context(), c.Params("id"))
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
