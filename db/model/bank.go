package model

type Bank struct {
	ID                string `gorm:"primary_key"`
	BusinessPartnerID string
	Name              string
	BranchName        string
	BankHolder        string
	Timestamp
}
