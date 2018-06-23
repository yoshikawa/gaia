package models

import "time"

// Sensor センサーに関わるモデル
type Sensor struct {
	ID             int64     `db:"id, primarykey, autoincrement" json:"id"`
	OrganizationID int64     `db:"organization_id" json:"organization_id"`
	Name           string    `db:"name" json:"name"`
	Email          string    `db:"email" json:"email"`
	Password       string    `db:"password" json:"password"`
	Country        string    `db:"country" json:"country"`
	Administrator  bool      `db:"administrator" json:"administrator"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
