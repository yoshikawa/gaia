package models

import "time"

// ObservationPosition 観測位置に関わるモデル
type ObservationPosition struct {
	ID        int64     `db:"id, primarykey, autoincrement" json:"id"`
	FieldID   int64     `db:"field_id" json:"field_id"`
	PlantID   int64     `db:"plant_id" json:"plant_id"`
	Name      string    `db:"name" json:"name"`
	Point     float64   `db:"point" json:"point"`
	Altitude  float64   `db:"altitude" json:"altitude"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Deleted   time.Time `db:"deleted" json:"deleted"`
}
