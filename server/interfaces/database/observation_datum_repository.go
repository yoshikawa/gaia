package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
)

// ObservationDatumRepository model
type ObservationDatumRepository struct {
	SQLHandler
}

// Store insert values into observation data table
func (repo *ObservationDatumRepository) Store(od domain.ObservationDatum) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO observation_data (
			observation_position_id,
			sensor_id,
			value,
			datetime
		) VALUES (
			?,
			?,
			?,
			?
		)`,
		od.ObservationPositionID,
		od.SensorID,
		od.Value,
		od.Datetime,
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

// FindByID find the observation datum by id
func (repo *ObservationDatumRepository) FindByID(identifier int64) (od domain.ObservationDatum, err error) {
	row, err := repo.Query(`SELECT * FROM observation_data WHERE id = ?`, identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var observationPositionID int64
	var sensorID int64
	var value float64
	var datetime time.Time
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &observationPositionID, &sensorID, &value, &datetime, &createdAt, &updatedAt); err != nil {
		return
	}
	od.ID = id
	od.ObservationPositionID = observationPositionID
	od.SensorID = sensorID
	od.Value = value
	od.Datetime = datetime
	od.CreatedAt = createdAt
	od.UpdatedAt = updatedAt
	return
}

// FindAll find all observation data
func (repo *ObservationDatumRepository) FindAll() (od domain.ObservationData, err error) {
	rows, err := repo.Query("SELECT * FROM observation_data")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var observationPositionID int64
		var sensorID int64
		var value float64
		var datetime time.Time
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &observationPositionID, &sensorID, &value, &datetime, &createdAt, &updatedAt); err != nil {
			continue
		}
		datum := domain.ObservationDatum{
			ID:                    id,
			ObservationPositionID: observationPositionID,
			SensorID:              sensorID,
			Value:                 value,
			Datetime:              datetime,
			CreatedAt:             createdAt,
			UpdatedAt:             updatedAt,
		}
		od = append(od, datum)
	}
	return
}
