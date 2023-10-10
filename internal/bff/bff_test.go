package bff

import (
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/renanbastos93/transaction-routine/internal/app"
	"github.com/renanbastos93/transaction-routine/mocks"
)

func TestGetAccounts(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockApp())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Get("/accounts", bff.getAllAccounts)

		resp, err := app.Test(httptest.NewRequest("GET", "/accounts", nil))
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
	})
}

func TestGetAccountsWithErrorNotFound(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockGetAllAccountsWithError())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Get("/accounts", bff.getAllAccounts)

		resp, err := app.Test(httptest.NewRequest("GET", "/accounts", nil))
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 404, resp.StatusCode, "Status code")
	})
}

func TestGetAccountById(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockApp())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Get("/accounts/:id", bff.getAccountById)

		resp, err := app.Test(httptest.NewRequest("GET", "/accounts/1", nil))
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
	})
}

func TestGetAccountByIdNotFound(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockGetUserByIdErrorNotFound())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Get("/accounts/:id", bff.getAccountById)

		resp, err := app.Test(httptest.NewRequest("GET", "/accounts/123", nil))
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 404, resp.StatusCode, "Status code")
	})
}

func TestCreateAccount(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockApp())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Post("/accounts", bff.createAccount)

		payload := `{"document_number": "9223737139"}`
		req := httptest.NewRequest("POST", "/accounts", strings.NewReader(payload))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 201, resp.StatusCode, "Status code")
	})
}

func TestGetOperations(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockApp())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Get("/operations", bff.getAllOperations)

		resp, err := app.Test(httptest.NewRequest("GET", "/operations", nil))
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
	})
}

func TestGetOperationsWithError(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockGetAllOperationsWithError())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Get("/operations", bff.getAllOperations)

		resp, err := app.Test(httptest.NewRequest("GET", "/operations", nil))
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 404, resp.StatusCode, "Status code")
	})
}

func TestCreateAccountBodyInvalid(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockApp())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Post("/accounts", bff.createAccount)

		req := httptest.NewRequest("POST", "/accounts", nil)
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 400, resp.StatusCode, "Status code")
	})
}

func TestCreateAccountInvalidCreated(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockCreateUserError())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Post("/accounts", bff.createAccount)

		payload := `{"document_number": "9223737139"}`
		req := httptest.NewRequest("POST", "/accounts", strings.NewReader(payload))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 500, resp.StatusCode, "Status code")
	})
}

func TestSaveTransaction(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockApp())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Post("/transactions", bff.saveTransaction)

		payload := `{ "account_id": "b7f68815-cf29-4ea5-9282-05623c3e030f", "operation_type_id": "0330bfe8-efc8-4ef8-b8fc-d2462df3a3c1", "amout": 5 }`
		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(payload))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 201, resp.StatusCode, "Status code")
	})
}

func TestSaveTransactionErrorInvalidBody(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockApp())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Post("/transactions", bff.saveTransaction)

		payload := `{ "account_id": "b7f68815-cf29-4ea5-9282-05623c3e030f", "operation_type_id": "0330bfe8-efc8-4ef8-b8fc-d2462df3a3c1", "amout": 5 }`
		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(payload))
		req.Header.Add(fiber.HeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 400, resp.StatusCode, "Status code")
	})
}

func TestSaveTransactionWithErrorGetUserForbidden(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockGetUserByIdErrorForbidden())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Post("/transactions", bff.saveTransaction)

		payload := `{ "account_id": "b7f68815-cf29-4ea5-9282-05623c3e030f", "operation_type_id": "0330bfe8-efc8-4ef8-b8fc-d2462df3a3c1", "amout": 5 }`
		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(payload))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 403, resp.StatusCode, "Status code")
	})
}

func TestSaveTransactionWithErrorOperationTypeInvalid(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockThereIsOperationWithError())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Post("/transactions", bff.saveTransaction)

		payload := `{ "account_id": "b7f68815-cf29-4ea5-9282-05623c3e030f", "operation_type_id": "0330bfe8-efc8-4ef8-b8fc-d2462df3a3c1", "amout": 5 }`
		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(payload))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 400, resp.StatusCode, "Status code")
	})
}

func TestSaveTransactionWithErrorToSave(t *testing.T) {
	runner := weavertest.Local
	appFake := weavertest.Fake[app.Component](mocks.NewMockSaveTransactionError())
	runner.Fakes = append(weavertest.Local.Fakes, appFake)
	runner.Test(t, func(t *testing.T, bff *implBFF) {
		app := fiber.New()
		defer app.Shutdown()
		app.Post("/transactions", bff.saveTransaction)

		payload := `{ "account_id": "b7f68815-cf29-4ea5-9282-05623c3e030f", "operation_type_id": "0330bfe8-efc8-4ef8-b8fc-d2462df3a3c1", "amout": 5 }`
		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(payload))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, 500, resp.StatusCode, "Status code")
	})
}
