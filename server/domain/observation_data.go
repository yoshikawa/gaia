package domain

import "time"

// ObservationData Model
type ObservationData struct {
	ID                    int64     `json:"id" db:"id"`
	ObservationPositionID int64     `json:"observation_position_id" db:"observation_position_id"`
	SensorID              int64     `json:"sensor_id" db:"sensor_id"`
	Value                 float64   `json:"value" db:"value"`
	Datetime              time.Time `json:"datetime" db:"datetime"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
}
