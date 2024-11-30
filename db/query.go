package db

import (
	"github.com/icchy-san/uctest/db/model"
	"gorm.io/gorm"
	"time"
)

// GetInvoices queries database and return data which is within passed period.
// period_start_at and period_end_at are optional value, hence they are nullable.
// the period_start_at is inclusive and the period_end_at is exclusive.
func (d *db) GetInvoices(tx *gorm.DB, period_start_at, period_end_at *time.Time) ([]model.Invoice, error) {
	client := d.getClient(tx)
	var invoices []model.Invoice

	if period_start_at != nil {
		client = client.Where("due_date >= ?", period_start_at)
	}

	if period_end_at != nil {
		client = client.Where("due_date < ?", period_end_at)
	}

	if err := client.
		Find(&invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}
