package models

import "time"

// Field Model
type User struct {
	Id        int64     `db:"id" json:"id"`
	Groups    string    `db:"groups" json:"groups"`
	Name      string    `db:"name" json:"name"`
	Area      []float64 `db:"area" json:"area"`
	createdAt time.Time `db:"created_at" json:"created_at"`
	updatedAt time.Time `db:"updated_at" json:"updated_at"`
	deleted   bool      `db:"deleted" json:"deleted"`
}
