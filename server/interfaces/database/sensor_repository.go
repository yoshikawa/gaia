package database

import (
	"time"

	"github.com/Pluslab/fieldsensing/server/domain"
)

// SensorRepository model
type SensorRepository struct {
	SQLHandler
}

// Store insert values into sensor table
func (repo *SensorRepository) Store(sensor domain.Sensor) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO sensors (
			sensing_device_id,
			observable_property_id,
			individual_id,
			sensor_number,
			observation_condition
		) VALUES (
			?,
			?,
			?,
			?,
			?
		)`,
		sensor.SensingDeviceID,
		sensor.ObservablePropertyID,
		sensor.IndividualID,
		sensor.SensorNumber,
		sensor.ObservationCondition,
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

// FindByID find the sensor by id
func (repo *SensorRepository) FindByID(identifier int64) (sensor domain.Sensor, err error) {
	row, err := repo.Query("SELECT * FROM sensors WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var sensingDeviceID int64
	var observablePropertyID int64
	var individualID int64
	var sensorNumber string
	var observationCondition float64
	var createdAt time.Time
	var updatedAt time.Time
	var deleted bool
	row.Next()
	if err = row.Scan(&id, &sensingDeviceID, &observablePropertyID, &individualID, &sensorNumber, &observationCondition, &createdAt, &updatedAt, &deleted); err != nil {
		return
	}
	sensor.ID = id
	sensor.SensingDeviceID = sensingDeviceID
	sensor.ObservablePropertyID = observablePropertyID
	sensor.IndividualID = individualID
	sensor.SensorNumber = sensorNumber
	sensor.ObservationCondition = observationCondition
	sensor.CreatedAt = createdAt
	sensor.UpdatedAt = updatedAt
	sensor.Deleted = deleted
	return
}

// FindAll find all sensors
func (repo *SensorRepository) FindAll() (sensors domain.Sensors, err error) {
	rows, err := repo.Query("SELECT * FROM sensors")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var sensingDeviceID int64
		var observablePropertyID int64
		var individualID int64
		var sensorNumber string
		var observationCondition float64
		var createdAt time.Time
		var updatedAt time.Time
		var deleted bool
		if err := rows.Scan(&id, &sensingDeviceID, &observablePropertyID, &individualID, &sensorNumber, &observationCondition, &createdAt, &updatedAt, &deleted); err != nil {
			continue
		}
		sensor := domain.Sensor{
			ID:                   id,
			SensingDeviceID:      sensingDeviceID,
			ObservablePropertyID: observablePropertyID,
			IndividualID:         individualID,
			SensorNumber:         sensorNumber,
			ObservationCondition: observationCondition,
			CreatedAt:            createdAt,
			UpdatedAt:            updatedAt,
			Deleted:              deleted,
		}
		sensors = append(sensors, sensor)
	}
	return
}
