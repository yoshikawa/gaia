package models

import "time"

// Field Model
type Field struct {
	ID        int64     `db:"id, autoincrement" json:"id"`
	Groups    string    `db:"groups" json:"groups"`
	Name      string    `db:"name" json:"name"`
	Area      []float64 `db:"area" json:"area"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Deleted   bool      `db:"deleted" json:"deleted"`
}
