package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
)

// SensingDeviceStatusRepository model
type SensingDeviceStatusRepository struct {
	SQLHandler
}

// Store insert values into field table
func (repo *SensingDeviceStatusRepository) Store(sensingDeviceStatus domain.SensingDeviceStatus) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO sensing_device_status (
			sensing_device_id,
			battery,
			status,
			datetime
		) VALUES (
			?,
			?,
			?,
			?
		)`,
		sensingDeviceStatus.SensingDeviceID,
		sensingDeviceStatus.Battery,
		sensingDeviceStatus.Status,
		sensingDeviceStatus.Datetime,
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

// FindByID find the field by id
func (repo *SensingDeviceStatusRepository) FindByID(identifier int64) (sensingDeviceStatus domain.SensingDeviceStatus, err error) {
	row, err := repo.Query("SELECT * FROM sensing_device_status WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var sensingDeviceID int64
	var battery int16
	var status int16
	var datetime time.Time
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &sensingDeviceID, &battery, &status, &datetime, &createdAt, &updatedAt); err != nil {
		return
	}
	sensingDeviceStatus.ID = id
	sensingDeviceStatus.SensingDeviceID = sensingDeviceID
	sensingDeviceStatus.Battery = battery
	sensingDeviceStatus.Status = status
	sensingDeviceStatus.Datetime = datetime
	sensingDeviceStatus.CreatedAt = createdAt
	sensingDeviceStatus.UpdatedAt = updatedAt
	return
}

// FindAll find all fields
func (repo *SensingDeviceStatusRepository) FindAll() (sensingDeviceStatuses domain.SensingDeviceStatuses, err error) {
	rows, err := repo.Query("SELECT * FROM sensing_device_status")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var sensingDeviceID int64
		var battery int16
		var status int16
		var datetime time.Time
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &sensingDeviceID, &battery, &status, &datetime, &createdAt, &updatedAt); err != nil {
			continue
		}
		sensingDeviceStatus := domain.SensingDeviceStatus{
			ID:              id,
			SensingDeviceID: sensingDeviceID,
			Battery:         battery,
			Status:          status,
			Datetime:        datetime,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
		sensingDeviceStatuses = append(sensingDeviceStatuses, sensingDeviceStatus)
	}
	return
}
