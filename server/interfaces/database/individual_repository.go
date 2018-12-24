package database

import (
	"time"

	"github.com/Pluslab/fieldsensing/server/domain"
)

// IndividualRepository model
type IndividualRepository struct {
	SQLHandler
}

// Store insert values into individual table
func (repo *IndividualRepository) Store(individual domain.Individual) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO individuals (
			short_name,
			long_name,
			model_number,
			intended_application,
			sensor_type,
			minimum_observation_range,
			maximum_observation_range,
			code_of_observation_range,
			minimum_observation_accuracy_of_measured_value,
			maximum_observation_accuracy_of_measured_value,
			code_of_observation_accuracy,
			observation_resolution_of_measured_value,
			code_of_observation_resolution,
			observation_interval_strict,
			observation_timing_strict,
			manufacturer,
			location_name
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)`,
		individual.ShortName,
		individual.LongName,
		individual.ModelNumber,
		individual.IntendedApplication,
		individual.SensorType,
		individual.MinimumObservationRange,
		individual.MaximumObservationRange,
		individual.CodeOfObservationAccuracy,
		individual.ObservationResolutionOfMeasuredValue,
		individual.CodeOfObservationResolution,
		individual.ObservationIntervalStrict,
		individual.ObservationTimingStrict,
		individual.Manufacturer,
		individual.LocationName,
		individual.CreatedAt,
		individual.UpdatedAt,
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

// FindByID find the individual by id
func (repo *IndividualRepository) FindByID(identifier int64) (individual domain.Individual, err error) {
	row, err := repo.Query(
		`SELECT
			id,
			short_name,
			long_name,
			model_number,
			intended_application,
			sensor_type,
			minimum_observation_range,
			maximum_observation_range,
			code_of_observation_range,
			minimum_observation_accuracy_of_measured_value,
			maximum_observation_accuracy_of_measured_value,
			code_of_observation_accuracy,
			observation_resolution_of_measured_value,
			code_of_observation_resolution,
			observation_interval_strict,
			observation_timing_strict,
			manufacturer,
			location_name
		FROM individuals WHERE id = ?`, identifier,
	)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var shortName string
	var longName string
	var modelNumber string
	var intendedApplication string
	var sensorType string
	var minimumObservationRange float64
	var maximumObservationRange float64
	var codeOfObservationRange string
	var minimumObservationAccuracyOfMeasuredValue float64
	var maximumObservationAccuracyOfMeasuredValue float64
	var codeOfObservationAccuracy string
	var observationResolutionOfMeasuredValue float64
	var codeOfObservationResolution string
	var observationIntervalStrict bool
	var observationTimingStrict bool
	var manufacturer string
	var locationName string
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(
		&id,
		&shortName,
		&longName,
		&modelNumber,
		&intendedApplication,
		&sensorType,
		&minimumObservationRange,
		&maximumObservationRange,
		&codeOfObservationRange,
		&minimumObservationAccuracyOfMeasuredValue,
		&maximumObservationAccuracyOfMeasuredValue,
		&codeOfObservationAccuracy,
		&observationResolutionOfMeasuredValue,
		&codeOfObservationResolution,
		&observationIntervalStrict,
		&observationTimingStrict,
		&manufacturer,
		&locationName,
		&createdAt,
		&updatedAt,
	); err != nil {
		return
	}
	individual.ID = id
	individual.ShortName = shortName
	individual.LongName = longName
	individual.ModelNumber = modelNumber
	individual.IntendedApplication = intendedApplication
	individual.SensorType = sensorType
	individual.MinimumObservationRange = minimumObservationRange
	individual.MaximumObservationRange = maximumObservationRange
	individual.CodeOfObservationRange = codeOfObservationRange
	individual.MinimumObservationAccuracyOfMeasuredValue = minimumObservationAccuracyOfMeasuredValue
	individual.MaximumObservationAccuracyOfMeasuredValue = maximumObservationAccuracyOfMeasuredValue
	individual.CodeOfObservationAccuracy = codeOfObservationAccuracy
	individual.ObservationResolutionOfMeasuredValue = observationResolutionOfMeasuredValue
	individual.CodeOfObservationResolution = codeOfObservationResolution
	individual.ObservationIntervalStrict = observationIntervalStrict
	individual.ObservationTimingStrict = observationTimingStrict
	individual.Manufacturer = manufacturer
	individual.LocationName = locationName
	individual.CreatedAt = createdAt
	individual.UpdatedAt = updatedAt
	return
}

// FindAll find all individuals
func (repo *IndividualRepository) FindAll() (individuals domain.Individuals, err error) {
	rows, err := repo.Query(`SELECT * FROM individuals`)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var shortName string
		var longName string
		var modelNumber string
		var intendedApplication string
		var sensorType string
		var minimumObservationRange float64
		var maximumObservationRange float64
		var codeOfObservationRange string
		var minimumObservationAccuracyOfMeasuredValue float64
		var maximumObservationAccuracyOfMeasuredValue float64
		var codeOfObservationAccuracy string
		var observationResolutionOfMeasuredValue float64
		var codeOfObservationResolution string
		var observationIntervalStrict bool
		var observationTimingStrict bool
		var manufacturer string
		var locationName string
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(
			&id,
			&shortName,
			&longName,
			&modelNumber,
			&intendedApplication,
			&sensorType,
			&minimumObservationRange,
			&maximumObservationRange,
			&codeOfObservationRange,
			&minimumObservationAccuracyOfMeasuredValue,
			&maximumObservationAccuracyOfMeasuredValue,
			&codeOfObservationAccuracy,
			&observationResolutionOfMeasuredValue,
			&codeOfObservationResolution,
			&observationIntervalStrict,
			&observationTimingStrict,
			&manufacturer,
			&locationName,
			&createdAt,
			&updatedAt,
		); err != nil {
			continue
		}
		individual := domain.Individual{
			ID:                      id,
			ShortName:               shortName,
			LongName:                longName,
			ModelNumber:             modelNumber,
			IntendedApplication:     intendedApplication,
			SensorType:              sensorType,
			MinimumObservationRange: minimumObservationRange,
			MaximumObservationRange: maximumObservationRange,
			CodeOfObservationRange:  codeOfObservationRange,
			MinimumObservationAccuracyOfMeasuredValue: minimumObservationAccuracyOfMeasuredValue,
			MaximumObservationAccuracyOfMeasuredValue: maximumObservationAccuracyOfMeasuredValue,
			CodeOfObservationAccuracy:                 codeOfObservationAccuracy,
			ObservationResolutionOfMeasuredValue:      observationResolutionOfMeasuredValue,
			CodeOfObservationResolution:               codeOfObservationResolution,
			ObservationIntervalStrict:                 observationIntervalStrict,
			ObservationTimingStrict:                   observationTimingStrict,
			Manufacturer:                              manufacturer,
			LocationName:                              locationName,
			CreatedAt:                                 createdAt,
			UpdatedAt:                                 updatedAt,
		}
		individuals = append(individuals, individual)
	}
	return
}
