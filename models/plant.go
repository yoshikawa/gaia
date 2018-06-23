package models

import "time"

// Plant 植物に関わるモデル
type Plant struct {
	ID                int64     `db:"id, primarykey, autoincrement" json:"id"`
	PlantCategoriesID int64     `db:"plant_categories_id" json:"plant_categories_id"`
	Variety           string    `db:"variety" json:"variety"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}
