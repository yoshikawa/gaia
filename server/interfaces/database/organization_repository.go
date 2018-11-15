package database

import "github.com/Pluslab/fieldsensing/server/domain"

// OrganizationRepository is organization's repository
type OrganizationRepository struct {
	SQLHandler
}

// Store insert values into organization table
func (repo *OrganizationRepository) Store(o domain.Organization) (id int, err error) {
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
	id = int(id64)
	return
}
