package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/icchy-san/uctest/gen/api"
)

type ServiceAPI struct{}

func (s *ServiceAPI) GetInvoices(ctx context.Context, req api.GetInvoicesRequestObject) (api.GetInvoicesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceAPI) PostInvoices(ctx context.Context, req api.PostInvoicesRequestObject) (api.PostInvoicesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func New(ctx context.Context) *fiber.App {
	serviceAPI := &ServiceAPI{}

	app := fiber.New(fiber.Config{})
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	apiGroup := app.Group("/api")

	server := api.NewStrictHandler(serviceAPI, nil)
	api.RegisterHandlers(apiGroup, server)

	return app
}
