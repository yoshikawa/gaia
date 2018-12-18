package database

import (
	"time"

	"github.com/Pluslab/fieldsensing/server/domain"
)

// GatewayDeviceRepository model
type GatewayDeviceRepository struct {
	SQLHandler
}

// Store insert values into gateway device table
func (repo *GatewayDeviceRepository) Store(gatewayDevice domain.GatewayDevice) (id int64, err error) {
	result, err := repo.Execute(
		"INSERT INTO gateway_devices (vendor_id, device_type, serial_number, point, name, transmission_distance, authentication_name, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		gatewayDevice.VendorID, gatewayDevice.DeviceType, gatewayDevice.SerialNumber,
		gatewayDevice.Point, gatewayDevice.Name, gatewayDevice.TransmissionDistance,
		gatewayDevice.AuthenticationName, gatewayDevice.Password,
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

// FindByID find the gateway device by id
func (repo *GatewayDeviceRepository) FindByID(identifier int64) (gatewayDevice domain.GatewayDevice, err error) {
	row, err := repo.Query("SELECT vendor_id, device_type, serial_number, point, name, transmission_distance, authentication_name, password FROM gateway_devices WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var vendorID int64
	var deviceType string
	var serialNumber string
	var point float64
	var name string
	var transmissionDistance int16
	var authenticationName string
	var password string
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &vendorID, &deviceType, &serialNumber, &point, &name, &transmissionDistance, &authenticationName, &password, &updatedAt); err != nil {
		return
	}
	gatewayDevice.ID = id
	gatewayDevice.VendorID = vendorID
	gatewayDevice.DeviceType = deviceType
	gatewayDevice.SerialNumber = serialNumber
	gatewayDevice.Point = point
	gatewayDevice.Name = name
	gatewayDevice.TransmissionDistance = transmissionDistance
	gatewayDevice.AuthenticationName = authenticationName
	gatewayDevice.Password = password
	gatewayDevice.CreatedAt = createdAt
	gatewayDevice.UpdatedAt = updatedAt
	return
}

// FindAll find all gateway devicess
func (repo *GatewayDeviceRepository) FindAll() (gatewayDevices domain.GatewayDevices, err error) {
	rows, err := repo.Query("SELECT id, organization_id, name, area, created_at, updated_at, deleted FROM gateway_devices")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var vendorID int64
		var deviceType string
		var serialNumber string
		var point float64
		var name string
		var transmissionDistance int16
		var authenticationName string
		var password string
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &vendorID, &deviceType, &serialNumber, &point, &name, &transmissionDistance, &authenticationName, &password, &updatedAt); err != nil {
			continue
		}
		gatewayDevice := domain.GatewayDevice{
			ID:                   id,
			VendorID:             vendorID,
			DeviceType:           deviceType,
			SerialNumber:         serialNumber,
			Point:                point,
			Name:                 name,
			TransmissionDistance: transmissionDistance,
			AuthenticationName:   authenticationName,
			Password:             password,
			CreatedAt:            createdAt,
			UpdatedAt:            updatedAt,
		}
		gatewayDevices = append(gatewayDevices, gatewayDevice)
	}
	return
}
