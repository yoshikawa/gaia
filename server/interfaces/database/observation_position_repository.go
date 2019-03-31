package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
)

// ObservationPositionRepository model
type ObservationPositionRepository struct {
	SQLHandler
}

// Store insert values into observation position table
func (repo *ObservationPositionRepository) Store(op domain.ObservationPosition) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO observation_positions (
			field_id,
			plant_id,
			name,
			point,
			altitude
		) VALUES (
			?,
			?,
			?,
			?,
			?)`,
		op.FieldID,
		op.PlantID,
		op.Name,
		op.Point,
		op.Altitude,
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

// FindByID find the observation position by id
func (repo *ObservationPositionRepository) FindByID(identifier int64) (op domain.ObservationPosition, err error) {
	row, err := repo.Query("SELECT * FROM observation_positions WHERE id = ? LIMIT 1", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var fieldID int64
	var plantID int64
	var name string
	var point float64
	var altitude float64
	var createdAt time.Time
	var updatedAt time.Time
	var deleted bool
	row.Next()
	if err = row.Scan(&id, &fieldID, &plantID, &name, &point, &altitude, &createdAt, &updatedAt, &deleted); err != nil {
		return
	}
	op.ID = id
	op.FieldID = fieldID
	op.PlantID = plantID
	op.Name = name
	op.Point = point
	op.Altitude = altitude
	op.CreatedAt = createdAt
	op.UpdatedAt = updatedAt
	op.Deleted = deleted
	return
}

// FindAll find all observation positions
func (repo *ObservationPositionRepository) FindAll() (ops domain.ObservationPositions, err error) {
	rows, err := repo.Query("SELECT * FROM observation_positions")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var fieldID int64
		var plantID int64
		var name string
		var point float64
		var altitude float64
		var createdAt time.Time
		var updatedAt time.Time
		var deleted bool
		if err := rows.Scan(&id, &fieldID, &plantID, &name, &point, &altitude, &createdAt, &updatedAt, &deleted); err != nil {
			continue
		}
		op := domain.ObservationPosition{
			ID:        id,
			FieldID:   fieldID,
			PlantID:   plantID,
			Name:      name,
			Point:     point,
			Altitude:  altitude,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Deleted:   deleted,
		}
		ops = append(ops, op)
	}
	return
}
