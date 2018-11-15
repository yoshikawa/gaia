package domain

import "time"

// Plant Model
type Plant struct {
	ID                int64     `json:"id" db:"id"`
	PlantCategoriesID int64     `json:"plant_categories_id" db:"plant_categories_id"`
	Variety           string    `json:"variety" db:"variety"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}
