package database

import "github.com/Pluslab/fieldsensing/server/domain"

// OrganizationRepository is organization's repository
type OrganizationRepository struct {
	SQLHandler
}

// Store insert values into organization table
func (repo *OrganizationRepository) Store(o domain.Organization) (id int64, err error) {
	result, err := repo.Execute(
		"INSERT INTO organization (name) VALUES (?)", o.Name,
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
func (repo *OrganizationRepository) FindByID(identifier int) (organization domain.Organization, err error) {
	row, err := repo.Query("SELECT id, name FROM organization WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var name string
	row.Next()
	if err = row.Scan(&id, &name); err != nil {
		return
	}
	organization.ID = id
	organization.Name = name
	return
}

// FindAll find all organizations
func (repo *OrganizationRepository) FindAll() (organizations domain.Organizations, err error) {
	rows, err := repo.Query("SELECT id, name FROM organizations")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		organization := domain.Organization{
			ID:   id,
			Name: name,
		}
		organizations = append(organizations, organization)
	}
	return
}
