package domain

import "time"

// ObservationPosition Model
type ObservationPosition struct {
	ID        int64     `json:"id" db:"id"`
	FieldID   int64     `json:"field_id" db:"field_id"`
	PlantID   int64     `json:"plant_id" db:"plant_id"`
	Name      string    `json:"name" db:"name"`
	Point     float64   `json:"point" db:"point"`
	Altitude  float64   `json:"altitude" db:"altitude"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Deleted   time.Time `json:"deleted" db:"deleted"`
}
