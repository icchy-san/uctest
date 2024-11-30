package model

import "time"

type InvoiceStatus string

const (
	InvoiceStatusUnknown    InvoiceStatus = "UNKNOWN"
	InvoiceStatusUnPaid     InvoiceStatus = "UNPAID"
	InvoiceStatusProcessing InvoiceStatus = "PROCESSING"
	InvoiceStatusPaid       InvoiceStatus = "PAID"
	InvoiceStatusFailed     InvoiceStatus = "FAILED"
)

type Invoice struct {
	ID              string  `gorm:"primary_key"`
	Company         Company `gorm:"foreignKey:PublishedBy"`
	PublishedBy     string  // This ID user publishes this invoice.
	PublishedAt     time.Time
	BusinessPartner BusinessPartner `gorm:"foreignKey:ReceivedBy"`
	ReceivedBy      string
	Commission      int
	CommissionRate  float32
	Tax             int
	TaxRate         float32
	PaymentAmount   int
	BilledAmount    int
	DueDate         time.Time
	Status          InvoiceStatus
	Timestamp
}

func (s *InvoiceStatus) ToString() string {
	if s == nil {
		return string(InvoiceStatusUnknown)
	}

	return string(*s)
}
