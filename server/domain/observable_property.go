package domain

import "time"

// ObservableProperty Model
type ObservableProperty struct {
	ID                       int64     `json:"id" db:"id"`
	ObservablePropertyNumber string    `json:"observable_property_number" db:"observable_property_number"`
	System                   string    `json:"system" db:"system"`
	Classification           string    `json:"classification" db:"classification"`
	ClassificationEn         string    `json:"classification_en" db:"classification_en"`
	ObservationProperty      string    `json:"observation_property" db:"observation_property"`
	ObservationPropertyEn    string    `json:"observation_property_en" db:"observation_property_en"`
	UnitOfMeasure            string    `json:"unit_of_measure" db:"unit_of_measure"`
	DisplayUnit              string    `json:"display_unit" db:"display_unit"`
	CreatedAt                time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                time.Time `json:"updated_at" db:"updated_at"`
}

// ObservableProperties Model
type ObservableProperties []ObservableProperty
