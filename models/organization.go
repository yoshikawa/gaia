package models

import "time"

// User Model
type Organization struct {
	Id        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	createdAt time.Time `db:"created_at" json:"created_at"`
	updatedAt time.Time `db:"updated_at" json:"updated_at"`
}
