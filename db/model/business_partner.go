package model

type BusinessPartner struct {
	ID                 string `gorm:"primaryKey"`
	CompanyID          string
	Name               string
	RepresentativeName string
	PhoneNumber        string
	ZipCode            string
	Address            string
	Bank               Bank
	Timestamp
}
