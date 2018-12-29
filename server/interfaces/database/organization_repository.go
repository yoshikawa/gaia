package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
)

// OrganizationRepository is organization's repository
type OrganizationRepository struct {
	SQLHandler
}

// Store insert values into organization table
func (repo *OrganizationRepository) Store(o domain.Organization) (id int64, err error) {
	result, err := repo.Execute(
		"INSERT INTO organizations (name) VALUES (?)", o.Name,
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

// FindByID find the organization by id
func (repo *OrganizationRepository) FindByID(identifier int64) (organization domain.Organization, err error) {
	row, err := repo.Query("SELECT * FROM organizations WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var name string
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
		return
	}
	organization.ID = id
	organization.Name = name
	organization.CreatedAt = createdAt
	organization.UpdatedAt = updatedAt
	return
}

// FindAll find all organizations
func (repo *OrganizationRepository) FindAll() (organizations domain.Organizations, err error) {
	rows, err := repo.Query("SELECT * FROM organizations")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var name string
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
			continue
		}
		organization := domain.Organization{
			ID:        id,
			Name:      name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		organizations = append(organizations, organization)
	}
	return
}
