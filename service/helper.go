package service

import (
	"github.com/icchy-san/uctest/db/model"
	"github.com/icchy-san/uctest/gen/api"
)

func convertInvoiceModelsToAPIInvoices(invoices []model.Invoice) []api.Invoice {
	apiInvoices := make([]api.Invoice, len(invoices))
	for i, invoice := range invoices {
		apiInvoices[i] = convertInvoiceModelToAPIInvoice(invoice)
	}

	return apiInvoices
}

func convertInvoiceModelToAPIInvoice(invoice model.Invoice) api.Invoice {
	return api.Invoice{
		BilledAmount:   invoice.BilledAmount,
		Commission:     invoice.Commission,
		CommissionRate: invoice.CommissionRate,
		DueDate:        invoice.DueDate,
		Id:             invoice.ID,
		PaymentAmount:  invoice.PaymentAmount,
		PublishedAt:    invoice.PublishedAt,
		Status:         invoice.Status.ToString(),
		Tax:            invoice.Tax,
		TaxRate:        invoice.TaxRate,
	}
}
