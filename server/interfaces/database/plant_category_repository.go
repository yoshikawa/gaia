package database

import (
	"time"

	"github.com/Pluslab/fieldsensing/server/domain"
)

// PlantCategoryRepository model
type PlantCategoryRepository struct {
	SQLHandler
}

// Store insert values into field table
func (repo *PlantCategoryRepository) Store(plantCategory domain.PlantCategory) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO plant_categories (
			large_classification,
			middle_classification,
			small_classification,
			thesaurus,
			harvest_site,
			attribute_item
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?
		)`,
		plantCategory.LargeClassification,
		plantCategory.MiddleClassification,
		plantCategory.SmallClassification,
		plantCategory.Thesaurus,
		plantCategory.HarvestSite,
		plantCategory.AttributeItem,
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
func (repo *PlantCategoryRepository) FindByID(identifier int64) (plantCategory domain.PlantCategory, err error) {
	row, err := repo.Query("SELECT * FROM plant_categories WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var largeClassification string
	var middleClassification string
	var smallClassification string
	var thesaurus string
	var harvestSite string
	var attributeItem string
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &largeClassification, &middleClassification, &smallClassification, &thesaurus, &harvestSite, &attributeItem, &createdAt, &updatedAt); err != nil {
		return
	}
	plantCategory.ID = id
	plantCategory.LargeClassification = largeClassification
	plantCategory.MiddleClassification = middleClassification
	plantCategory.SmallClassification = smallClassification
	plantCategory.Thesaurus = thesaurus
	plantCategory.HarvestSite = harvestSite
	plantCategory.AttributeItem = attributeItem
	plantCategory.CreatedAt = createdAt
	plantCategory.UpdatedAt = updatedAt
	return
}

// FindAll find all fields
func (repo *PlantCategoryRepository) FindAll() (plantCategories domain.PlantCategories, err error) {
	rows, err := repo.Query("SELECT id, organization_id, name, area, created_at, updated_at, deleted FROM fileds")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var largeClassification string
		var middleClassification string
		var smallClassification string
		var thesaurus string
		var harvestSite string
		var attributeItem string
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &largeClassification, &middleClassification, &smallClassification, &thesaurus, &harvestSite, &attributeItem, &createdAt, &updatedAt); err != nil {
			continue
		}
		plantCategory := domain.PlantCategory{
			ID:                   id,
			LargeClassification:  largeClassification,
			MiddleClassification: middleClassification,
			SmallClassification:  smallClassification,
			Thesaurus:            thesaurus,
			HarvestSite:          harvestSite,
			AttributeItem:        attributeItem,
			CreatedAt:            createdAt,
			UpdatedAt:            updatedAt,
		}
		plantCategories = append(plantCategories, plantCategory)
	}
	return
}
