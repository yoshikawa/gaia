package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
)

// PlantRepository model
type PlantRepository struct {
	SQLHandler
}

// Store insert values into field table
func (repo *PlantRepository) Store(plant domain.Plant) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO plants (
			plant_categories_id,
			variety
		) VALUES (
			?,
			?
		)`,
		plant.PlantCategoriesID,
		plant.Variety,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int64(id64)
	return
}

// FindByID find the field by id
func (repo *PlantRepository) FindByID(identifier int64) (plant domain.Plant, err error) {
	row, err := repo.Query(`SELECT * FROM fields WHERE id = ?`, identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var plantCategoriesID int64
	var variety string
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &plantCategoriesID, &variety, &createdAt, &updatedAt); err != nil {
		return
	}
	plant.ID = id
	plant.PlantCategoriesID = plantCategoriesID
	plant.Variety = variety
	plant.CreatedAt = createdAt
	plant.UpdatedAt = updatedAt
	return
}

// FindAll find all fields
func (repo *PlantRepository) FindAll() (plants domain.Plants, err error) {
	rows, err := repo.Query("SELECT id, organization_id, name, area, created_at, updated_at, deleted FROM fileds")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var plantCategoriesID int64
		var variety string
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &plantCategoriesID, &variety, &createdAt, &updatedAt); err != nil {
			continue
		}
		plant := domain.Plant{
			ID:                id,
			PlantCategoriesID: plantCategoriesID,
			Variety:           variety,
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
		}
		plants = append(plants, plant)
	}
	return
}
