package server

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/icchy-san/uctest/gen/api"
	"github.com/icchy-san/uctest/service"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

type ServiceAPI struct {
	invoiceService service.InvoiceService
}

func (s *ServiceAPI) GetInvoices(ctx context.Context, req api.GetInvoicesRequestObject) (api.GetInvoicesResponseObject, error) {
	return s.invoiceService.GetInvoices(ctx, req)
}

func (s *ServiceAPI) PostInvoices(ctx context.Context, req api.PostInvoicesRequestObject) (api.PostInvoicesResponseObject, error) {
	return s.invoiceService.PostInvoices(ctx, req)
}

func New(ctx context.Context, invoiceService service.InvoiceService) *fiber.App {
	serviceAPI := &ServiceAPI{
		invoiceService: invoiceService,
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			err = ctx.Status(code).JSON(ErrorResponse{
				Error: err.Error(),
				Code:  code,
			})
			if err != nil {
				return err
			}

			return nil
		},
	})
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	apiGroup := app.Group("/api")

	server := api.NewStrictHandler(serviceAPI, nil)
	api.RegisterHandlers(apiGroup, server)

	return app
}
