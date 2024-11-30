package service

import (
	"context"
	"fmt"
	"github.com/icchy-san/uctest/gen/api"
)

func (s *invoiceService) GetInvoices(ctx context.Context, req api.GetInvoicesRequestObject) (api.GetInvoicesResponseObject, error) {
	// precondition: retrieved period, start and end as optional values from query.
	periodStartAt := req.Params.PeriodStartAt
	periodEndAt := req.Params.PeriodEndAt

	if periodStartAt != nil && periodEndAt != nil {
		if periodStartAt.After(*periodEndAt) {
			return nil, fmt.Errorf("period_start_at %v is needed be before period_end_at %v", *periodStartAt, *periodEndAt)
		}
	}

	// get invoices data from db by using such periods.
	invoices, err := s.database.GetInvoices(nil, periodStartAt, periodEndAt)
	if err != nil {
		return nil, err
	}

	// convert model to response type
	var resInvoice api.GetInvoices200JSONResponse
	resInvoice = convertInvoiceModelsToAPIInvoices(invoices)

	return resInvoice, nil
}
