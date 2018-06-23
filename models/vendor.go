package models

import "time"

// Vendor ベンダーに関わるモデル
type Vendor struct {
	ID                    string    `db:"id, primarykey, autoincrement" json:"id"`
	Name                  string    `db:"name, type=varchar, size=255" json:"mame"`
	OrganizationName      string    `db:"organization_name" json:"organization_name"`
	Voice                 string    `db:"voice" json:"voice"`
	Facsimile             string    `db:"facsimile" json:"facsimile"`
	DelivaryPoint         string    `db:"delivary_point" json:"delivary_point"`
	City                  string    `db:"city" json:"city"`
	PostalCode            string    `db:"postal_code" json:"postal_code"`
	Country               string    `db:"country" json:"country"`
	ElectronicMailAddress string    `db:"electronic_mail_address" json:"electronic_mail_address"`
	CreatedAt             time.Time `db:"created_at" json:"created_at"`
	UpdatedAt             time.Time `db:"updated_at" json:"updated_at"`
}
