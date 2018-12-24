package domain

import "time"

// SensingDevice Model
type SensingDevice struct {
	ID                    int64     `json:"id" db:"id"`
	ObservationPositionID int64     `json:"observation_position_id" db:"observation_position_id"`
	VendorID              string    `json:"vendor_id" db:"vendor_id"`
	DeviceType            string    `json:"device_type" db:"device_type"`
	SerialNumber          string    `json:"serial_number" db:"serial_number"`
	Name                  string    `json:"name" db:"name"`
	TransmissionDistance  int16     `json:"transmission_distance" db:"transmission_distance"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
	Deleted               bool      `json:"deleted" db:"deleted"`
}

// SensingDevices Model
type SensingDevices []SensingDevice
