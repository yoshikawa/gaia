package domain

import "time"

// GatewayDeviceStatus Model
type GatewayDeviceStatus struct {
	ID              int64     `json:"id" db:"id"`
	GatewayDeviceID int64     `json:"gateway_device_id" db:"gateway_device_id"`
	Battery         int16     `json:"battery" db:"battery"`
	Status          int16     `json:"status" db:"status"`
	Datetime        time.Time `json:"datetime" db:"datetime"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// GatewayDeviceStatuses Model
type GatewayDeviceStatuses []GatewayDeviceStatus
