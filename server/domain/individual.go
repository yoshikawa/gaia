package domain

import (
	"time"
)

// Individual Model
type Individual struct {
	ID                                        int64     `json:"id" db:"id"`
	ShortName                                 string    `json:"short_name" db:"short_name"`
	LongName                                  string    `json:"long_name" db:"long_name"`
	ModelNumber                               string    `json:"model_number" db:"model_number"`
	IntendedApplication                       string    `json:"intended_application" db:"intended_application"`
	SensorType                                string    `json:"sensor_type" db:"sensor_type"`
	MinimumObservationRange                   float64   `json:"minimum_observation_range" db:"minimum_observation_range"`
	MaximumObservationRange                   float64   `json:"maximum_observation_range" db:"maximum_observation_range"`
	CodeOfObservationRange                    string    `json:"code_of_observation_range" db:"code_of_observation_range"`
	MinimumObservationAccuracyOfMeasuredValue float64   `json:"minimum_observation_accuracy_of_measured_value" db:"minimum_observation_accuracy_of_measured_value"`
	MaximumObservationAccuracyOfMeasuredValue float64   `json:"maximum_observation_accuracy_of_measured_value" db:"maximum_observation_accuracy_of_measured_value"`
	CodeOfObservationAccuracy                 string    `json:"code_of_observation_accuracy" db:"code_of_observation_accuracy"`
	ObservationResolutionOfMeasuredValue      float64   `json:"observation_resolution_of_measured_value" db:"observation_resolution_of_measured_value"`
	CodeOfObservationResolution               string    `json:"code_of_observation_resolution" db:"code_of_observation_resolution"`
	ObservationIntervalStrict                 bool      `json:"observation_interval_strict" db:"observation_interval_strict"`
	ObservationTimingStrict                   bool      `json:"observation_timing_strict" db:"observation_timing_strict"`
	Manufacturer                              string    `json:"manufacturer" db:"manufacturer"`
	LocationName                              string    `json:"location_name" db:"location_name"`
	CreatedAt                                 time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                                 time.Time `json:"updated_at" db:"updated_at"`
}
