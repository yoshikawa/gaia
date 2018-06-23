package models

import "time"

// ObservationData 観測データに関わるモデル
type ObservationData struct {
	ID                    int64     `db:"id, primarykey, autoincrement" json:"id"`
	ObservationPositionID int64     `db:"observation_position_id" json:"observation_position_id"`
	SensorID              int64     `db:"sensor_id" json:"sensor_id"`
	Value                 float64   `db:"value" json:"value"`
	Datetime              time.Time `db:"datetime" json:"datetime"`
	CreatedAt             time.Time `db:"created_at" json:"created_at"`
	UpdatedAt             time.Time `db:"updated_at" json:"updated_at"`
}
