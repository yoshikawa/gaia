package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
)

// GatewayDeviceStatusRepository model
type GatewayDeviceStatusRepository struct {
	SQLHandler
}

// Store insert values into gateway device table
func (repo *GatewayDeviceStatusRepository) Store(gatewayDeviceStatus domain.GatewayDeviceStatus) (id int64, err error) {
	result, err := repo.Execute(
		"INSERT INTO gateway_device_status (gateway_device_id, battery, status, datetime) VALUES (?, ?, ?, ?)",
		gatewayDeviceStatus.GatewayDeviceID, gatewayDeviceStatus.Battery, gatewayDeviceStatus.Status, gatewayDeviceStatus.Datetime,
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
func (repo *GatewayDeviceStatusRepository) FindByID(identifier int64) (gatewayDeviceStatus domain.GatewayDeviceStatus, err error) {
	row, err := repo.Query("SELECT gateway_device_id, battery, status, datetime FROM gateway_device_status WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var gatewayDeviceID int64
	var battery int16
	var status int16
	var datetime time.Time
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &gatewayDeviceID, &battery, &status, &datetime, &createdAt, &updatedAt); err != nil {
		return
	}
	gatewayDeviceStatus.ID = id
	gatewayDeviceStatus.GatewayDeviceID = gatewayDeviceID
	gatewayDeviceStatus.Battery = battery
	gatewayDeviceStatus.Status = status
	gatewayDeviceStatus.Datetime = datetime
	gatewayDeviceStatus.CreatedAt = createdAt
	gatewayDeviceStatus.UpdatedAt = updatedAt
	return
}

// FindAll find all gateway devicess
func (repo *GatewayDeviceStatusRepository) FindAll() (gatewayDeviceStatuses domain.GatewayDeviceStatuses, err error) {
	rows, err := repo.Query("SELECT id, vendor_id, device_type, serial_number, point, name, transmission_distance, authentication_name, password FROM gateway_devices")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var gatewayDeviceID int64
		var battery int16
		var status int16
		var datetime time.Time
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &gatewayDeviceID, &battery, &status, &datetime, &createdAt, &updatedAt); err != nil {
			continue
		}
		gatewayDeviceStatus := domain.GatewayDeviceStatus{
			ID:              id,
			GatewayDeviceID: gatewayDeviceID,
			Battery:         battery,
			Status:          status,
			Datetime:        datetime,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
		gatewayDeviceStatuses = append(gatewayDeviceStatuses, gatewayDeviceStatus)
	}
	return
}
