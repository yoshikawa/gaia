package domain

// User Model
type User struct {
	ID             int64  `json:"id" db:"id"`
	organizationID int64  `json:"organization_id" db:"organization_id"`
	Name           string `json:"name" db:"name"`
}
