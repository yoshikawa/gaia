package domain

import "time"

// Vendor model
type Vendor struct {
	ID                    int64     `json:"id" db:"id"`
	Name                  string    `json:"name" db:"name"`
	OrganizationName      string    `json:"organization_name" db:"organization_name"`
	Voice                 string    `json:"voice" db:"voice"`
	Facsimile             string    `json:"facsimile" db:"facsimile"`
	DelivaryPoint         string    `json:"delivary_point" db:"delivary_point"`
	City                  string    `json:"city" db:"city"`
	PostalCode            string    `json:"postal_code" db:"postal_code"`
	Country               string    `json:"country" db:"country"`
	ElectronicMailAddress string    `json:"electronic_mail_address" db:"electronic_mail_address"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
}

// Vendors model
type Vendors []Vendor
