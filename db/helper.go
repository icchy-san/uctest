package db

import "gorm.io/gorm"

func (d *db) getClient(tx *gorm.DB) *gorm.DB {
	client := d.client
	if tx != nil {
		client = tx
	}

	return client
}
