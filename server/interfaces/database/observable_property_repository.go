package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
)

// ObservablePropertyRepository model
type ObservablePropertyRepository struct {
	SQLHandler
}

// Store insert values into observable property table
func (repo *ObservablePropertyRepository) Store(observableProperty domain.ObservableProperty) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO observable_properties (
			observable_property_number,
			system,
			classification,
			classfication_en,
			observation_property,
			observation_property_en,
			unit_of_measure,
			display_unit
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)`,
		observableProperty.ObservablePropertyNumber,
		observableProperty.System,
		observableProperty.Classification,
		observableProperty.ClassificationEn,
		observableProperty.ObservationProperty,
		observableProperty.ObservationPropertyEn,
		observableProperty.UnitOfMeasure,
		observableProperty.DisplayUnit,
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

// FindByID find the observable property by id
func (repo *ObservablePropertyRepository) FindByID(identifier int64) (observableProperty domain.ObservableProperty, err error) {
	row, err := repo.Query("SELECT * FROM observable_properties WHERE id = ? LIMIT 1", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var observablePropertyNumber string
	var system string
	var classification string
	var classificationEn string
	var observationProperty string
	var observationPropertyEn string
	var unitOfMeasure string
	var displayUnit string
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(
		&id,
		&observablePropertyNumber,
		&system,
		&classification,
		&classificationEn,
		&observationProperty,
		&observationPropertyEn,
		&unitOfMeasure,
		&displayUnit,
		&createdAt,
		&updatedAt,
	); err != nil {
		return
	}
	observableProperty.ID = id
	observableProperty.ObservablePropertyNumber = observablePropertyNumber
	observableProperty.System = system
	observableProperty.Classification = classification
	observableProperty.ClassificationEn = classificationEn
	observableProperty.ObservationProperty = observationProperty
	observableProperty.ObservationPropertyEn = observationPropertyEn
	observableProperty.UnitOfMeasure = unitOfMeasure
	observableProperty.CreatedAt = createdAt
	observableProperty.UpdatedAt = updatedAt
	return
}

// FindAll find all observable properties
func (repo *ObservablePropertyRepository) FindAll() (observableProperties domain.ObservableProperties, err error) {
	rows, err := repo.Query("SELECT * FROM observable_properties")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var observablePropertyNumber string
		var system string
		var classification string
		var classificationEn string
		var observationProperty string
		var observationPropertyEn string
		var unitOfMeasure string
		var displayUnit string
		var createdAt time.Time
		var updatedAt time.Time
		if err = rows.Scan(
			&id,
			&observablePropertyNumber,
			&system,
			&classification,
			&classificationEn,
			&observationProperty,
			&observationPropertyEn,
			&unitOfMeasure,
			&displayUnit,
			&createdAt,
			&updatedAt,
		); err != nil {
			return
		}
		observableProperty := domain.ObservableProperty{
			ID:                       id,
			ObservablePropertyNumber: observablePropertyNumber,
			System:                   system,
			Classification:           classification,
			ClassificationEn:         classificationEn,
			ObservationProperty:      observationProperty,
			ObservationPropertyEn:    observationPropertyEn,
			UnitOfMeasure:            unitOfMeasure,
			DisplayUnit:              displayUnit,
			CreatedAt:                createdAt,
			UpdatedAt:                updatedAt,
		}
		observableProperties = append(observableProperties, observableProperty)
	}
	return
}
