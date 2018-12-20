package database

import (
	"time"

	"github.com/Pluslab/fieldsensing/server/domain"
)

// SensingDeviceRepository model
type SensingDeviceRepository struct {
	SQLHandler
}

// Store insert values into field table
func (repo *SensingDeviceRepository) Store(f domain.SensingDevice) (id int64, err error) {
	result, err := repo.Execute(
		"INSERT INTO fields (organization_id, name, area) VALUES (?, ?, ?)", f.OrganizationID, f.Name, f.Area,
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
func (repo *SensingDeviceRepository) FindByID(identifier int64) (field domain.SensingDevice, err error) {
	row, err := repo.Query("SELECT id, organization_id, name, area, created_at, updated_at, deleted FROM fields WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var organizationID int64
	var name string
	var area []float64
	var createdAt time.Time
	var updatedAt time.Time
	var deleted bool
	row.Next()
	if err = row.Scan(&id, &organizationID, &name, &area, &createdAt, &updatedAt, &deleted); err != nil {
		return
	}
	field.ID = id
	field.OrganizationID = organizationID
	field.Name = name
	field.Area = area
	field.CreatedAt = createdAt
	field.UpdatedAt = updatedAt
	field.Deleted = deleted
	return
}

// FindAll find all fields
func (repo *SensingDeviceRepository) FindAll() (fields domain.SensingDevices, err error) {
	rows, err := repo.Query("SELECT id, organization_id, name, area, created_at, updated_at, deleted FROM fileds")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var organizationID int64
		var name string
		var area []float64
		var createdAt time.Time
		var updatedAt time.Time
		var deleted bool
		if err := rows.Scan(&id, &organizationID, &name, &area, &createdAt, &updatedAt, &deleted); err != nil {
			continue
		}
		field := domain.SensingDevice{
			ID:             id,
			OrganizationID: organizationID,
			Name:           name,
			Area:           area,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			Deleted:        deleted,
		}
		fields = append(fields, field)
	}
	return
}
