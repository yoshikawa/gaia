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
func (organizationRepositoryGorm *OrganizationRepositoryGorm) Save(organization *domain.Organization) shared.Output {
	err := organizationRepositoryGorm.db.Save(organization).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: organization}
}

// Delete function
func (organizationRepositoryGorm *OrganizationRepositoryGorm) Delete(organization *domain.Organization) shared.Output {
	err := organizationRepositoryGorm.db.Delete(organization).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: organization}
}

// FindByID function
func (organizationRepositoryGorm *OrganizationRepositoryGorm) FindByID(id uint) shared.Output {
	var organization domain.Organization

	err := organizationRepositoryGorm.db.Where(&domain.Organization{ID: id}).Take(&organization).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: &organization}
}

// FindAll function
func (organizationRepositoryGorm *OrganizationRepositoryGorm) FindAll(params *shared.Parameters) shared.Output {
	var organizations domain.Organizations

	err := organizationRepositoryGorm.db.Offset(params.Offset).Limit(params.Limit).Order(params.OrderBy).Find(&organizations).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: organizations}
}

// Count function
func (organizationRepositoryGorm *OrganizationRepositoryGorm) Count(params *shared.Parameters) shared.Output {
	var count int
	err := organizationRepositoryGorm.db.Model(&domain.Organization{}).Offset(params.Offset).Limit(params.Limit).Order(params.OrderBy).Count(&count).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: count}
}
