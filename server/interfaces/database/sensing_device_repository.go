package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
)

// SensingDeviceRepository model
type SensingDeviceRepository struct {
	SQLHandler
}

// Store insert values into sensing device table
func (repo *SensingDeviceRepository) Store(sensingDevice domain.SensingDevice) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO sensing_devices (
			observation_position_id,
			vendor_id,
			device_type,
			serial_number,
			name,
			transmission_distance
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?
		)`,
		sensingDevice.ObservationPositionID,
		sensingDevice.VendorID,
		sensingDevice.DeviceType,
		sensingDevice.SerialNumber,
		sensingDevice.Name,
		sensingDevice.TransmissionDistance,
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

// FindByID find the sensing device by id
func (repo *SensingDeviceRepository) FindByID(identifier int64) (sensingDevice domain.SensingDevice, err error) {
	row, err := repo.Query("SELECT * FROM sensing_devices WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var observationPositionID int64
	var vendorID string
	var deviceType string
	var serialNumber string
	var name string
	var transmissionDistance int16
	var createdAt time.Time
	var updatedAt time.Time
	var deleted bool
	row.Next()
	if err = row.Scan(&id, &observationPositionID, &vendorID, &deviceType, &serialNumber, &name, &transmissionDistance, &createdAt, &updatedAt, &deleted); err != nil {
		return
	}
	sensingDevice.ID = id
	sensingDevice.ObservationPositionID = observationPositionID
	sensingDevice.VendorID = vendorID
	sensingDevice.DeviceType = deviceType
	sensingDevice.SerialNumber = serialNumber
	sensingDevice.Name = name
	sensingDevice.TransmissionDistance = transmissionDistance
	sensingDevice.CreatedAt = createdAt
	sensingDevice.UpdatedAt = updatedAt
	sensingDevice.Deleted = deleted
	return
}

// FindAll find all sensing devices
func (repo *SensingDeviceRepository) FindAll() (sensingDevices domain.SensingDevices, err error) {
	rows, err := repo.Query("SELECT * FROM sensing_devices")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var observationPositionID int64
		var vendorID string
		var deviceType string
		var serialNumber string
		var name string
		var transmissionDistance int16
		var createdAt time.Time
		var updatedAt time.Time
		var deleted bool
		if err := rows.Scan(&id, &observationPositionID, &vendorID, &deviceType, &serialNumber, &name, &transmissionDistance, &createdAt, &updatedAt, &deleted); err != nil {
			continue
		}
		sensingDevice := domain.SensingDevice{
			ID:                    id,
			ObservationPositionID: observationPositionID,
			VendorID:              vendorID,
			DeviceType:            deviceType,
			SerialNumber:          serialNumber,
			Name:                  name,
			TransmissionDistance:  transmissionDistance,
			CreatedAt:             createdAt,
			UpdatedAt:             updatedAt,
			Deleted:               deleted,
		}
		sensingDevices = append(sensingDevices, sensingDevice)
	}
	return
}
