package domain

import "time"

// Sensor Model
type Sensor struct {
	ID                   int64     `json:"id" db:"id"`
	SensingDeviceID      int64     `json:"sensing_device_id" db:"sensing_device_id"`
	ObservablePropertyID int64     `json:"observable_property_id" db:"observable_property_id"`
	IndividualID         int64     `json:"individual_id" db:"individual_id"`
	SensorNumber         string    `json:"sensor_number" db:"sensor_number"`
	ObservationCondition float64   `json:"observation_condition" db:"observation_condition"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
	Deleted              time.Time `json:"deleted" db:"deleted"`
}
