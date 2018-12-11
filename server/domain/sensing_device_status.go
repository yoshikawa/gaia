package domain

import "time"

// SensingDeviceStatus Model
type SensingDeviceStatus struct {
	ID              int64     `json:"id" db:"id"`
	SensingDeviceID int64     `json:"sensing_device_id" db:"sensing_device_id"`
	Battery         bool      `json:"battery" db:"battery"`
	Status          bool      `json:"status" db:"status"`
	Datetime        time.Time `json:"datetime" db:"datetime"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// SensingDeviceStatuses Model
type SensingDeviceStatuses []SensingDeviceStatus
