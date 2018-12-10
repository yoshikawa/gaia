package domain

import "time"

// Field Model
type Field struct {
	ID             int64     `json:"id" db:"id"`
	OrganizationID int64     `json:"organization_id" db:"organization_id" `
	Name           string    `json:"name" db:"name"`
	Area           []float64 `json:"area" db:"area"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Deleted        bool      `json:"deleted" db:"deleted"`
}

// Fields Model
type Fields []Field
