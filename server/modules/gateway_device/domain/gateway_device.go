package domain

import "time"

// GatewayDevice Model
type GatewayDevice struct {
	ID                   int64     `json:"id" db:"id"`
	VendorID             int64     `json:"sensing_device_id" db:"vendor_id"`
	DeviceType           string    `json:"device_type" db:"device_type"`
	SerialNumber         string    `json:"serial_number" db:"serial_number"`
	Point                float64   `json:"point" db:"point"`
	Name                 string    `json:"name" db:"name"`
	TransmissionDistance int16     `json:"transmission_distance" db:"transmission_distance"`
	AuthenticationName   string    `json:"authentication_name" db:"authentication_name"`
	Password             string    `json:"-" db:"password"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}

// GatewayDevices Model
type GatewayDevices []GatewayDevice
