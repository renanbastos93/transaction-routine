package bff

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/renanbastos93/transaction-routine/internal/app"
)

type implBFF struct {
	weaver.Implements[weaver.Main]
	app weaver.Ref[app.Component]
	bff weaver.Listener `weaver:"bff"`

	f *fiber.App
}

func (e *implBFF) createRouter(ctx context.Context) {
	log := logger.New(logger.ConfigDefault)
	// TODO: add others middleware to ensure security in this service

	router := e.f.Use(log)
	e.RegisterAccounts(router)
}

func Server(ctx context.Context, e *implBFF) (err error) {
	fmt.Printf("BFF listener available on %v\n", e.bff)

	e.f = fiber.New()
	e.createRouter(ctx)
	return e.f.Listener(e.bff)
}
