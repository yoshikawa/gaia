package repository

import (
	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/jinzhu/gorm"
)

// OrganizationRepositoryGorm struct
type OrganizationRepositoryGorm struct {
	db *gorm.DB
}

// NewOrganizationRepositoryGorm function
func NewOrganizationRepositoryGorm(db *gorm.DB) *OrganizationRepositoryGorm {
	return &OrganizationRepositoryGorm{db: db}
}

// Save function
func (r *OrganizationRepositoryGorm) Save(organization *domain.Organization) shared.Output {
	err := r.db.Save(organization).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: organization}
}
