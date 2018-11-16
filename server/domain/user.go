package domain

import "time"

// User Model
type User struct {
	ID             int64     `json:"id" db:"id"`
	OrganizationID int64     `json:"organization_id" db:"organization_id"`
	Name           string    `json:"name" db:"name"`
	Email          string    `json:"email" db:"email"`
	Password       string    `json:"password" db:"password"`
	Country        string    `json:"country" db:"country"`
	Administrator  bool      `json:"administrator" db:"administrator"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
