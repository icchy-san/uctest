package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/icchy-san/uctest/db/model"
	"github.com/icchy-san/uctest/gen/api"
)

func (s *invoiceService) PostInvoices(ctx context.Context, req api.PostInvoicesRequestObject) (api.PostInvoicesResponseObject, error) {
	// TODO: validate whether the received fields are correct or not.
	//   Like: Checking the company that publishes the invoice exits
	// calculate billed payment by using tax rate, commission rate, payment amount.
	// calculate formula is this: |
	//   支払金額 に手数料4%を加えたものに更に手数料の消費税を加えたものを請求金額とする
	//   ex: 10,000 + (10,000 * 0.04 * 1.10) = 10,440
	billedPayment, commission, tax := retrieveBilledPaymentWithTaxAndCommission(float32(req.Body.PaymentAmount), model.CommissionRate, model.TaxRate)

	idHash, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	invoiceID := fmt.Sprintf("in_%s", strings.ReplaceAll(idHash.String(), "-", ""))

	currentTime := time.Now()
	invoiceModel := &model.Invoice{
		ID:             invoiceID,
		PublishedBy:    req.Body.Publisher,
		PublishedAt:    currentTime,
		ReceivedBy:     req.Body.Receiver,
		Commission:     commission,
		CommissionRate: model.CommissionRate,
		Tax:            tax,
		TaxRate:        model.TaxRate,
		PaymentAmount:  req.Body.PaymentAmount,
		BilledAmount:   billedPayment,
		DueDate:        currentTime.AddDate(0, 1, 0),
		Status:         model.InvoiceStatusUnPaid,
	}

	if err := s.database.CreateInvoice(nil, invoiceModel); err != nil {
		return nil, err
	}

	return api.PostInvoices200JSONResponse(convertInvoiceModelToAPIInvoice(*invoiceModel)), nil
}

func retrieveBilledPaymentWithTaxAndCommission(paymentAmount, commissionRate, taxRate float32) (int, int, int) {
	// paymentAmount + (paymentAmount * commissionRate * taxRate)
	// ex: 10,000 + (10,000 * 0.04 * 1.10) = 10,440
	commission := paymentAmount * (commissionRate / 100) // 400 YEN
	totalFee := commission * (1 + (taxRate / 100))       // 440 YEN
	billedPayment := paymentAmount + totalFee            // 10000 + 440 YEN
	tax := totalFee - commission                         // 440 - 400 YEN

	return int(billedPayment), int(commission), int(tax)
}
