package models

import "time"

// SensingDevice センシングデバイスのモデル
type SensingDevice struct {
	ID              int64     `db:"id, primarykey, autoincrement" json:"id"`
	SensingDeviceID int64     `db:"sensing_device_id" json:"sensing_device_id"`
	Battery         bool      `db:"battery" json:"battery"`
	Status          bool      `db:"status" json:"status"`
	Datetime        time.Time `db:"datetime" json:"datetime"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
