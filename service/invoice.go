package service

import (
	"context"
	"github.com/icchy-san/uctest/gen/api"
)

type InvoiceService interface {
	GetInvoices(ctx context.Context, req api.GetInvoicesRequestObject) (api.GetInvoicesResponseObject, error)
	PostInvoices(ctx context.Context, req api.PostInvoicesRequestObject) (api.PostInvoicesResponseObject, error)
}

type invoiceService struct {
}

func New() InvoiceService {
	return &invoiceService{}
}
