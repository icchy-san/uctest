package model

type Company struct {
	ID                 string `gorm:"primaryKey"`
	CompanyID          string
	Name               string
	RepresentativeName string
	PhoneNumber        string
	ZipCode            string
	Address            string
	BusinessPartners   []BusinessPartner
	Timestamp
}
