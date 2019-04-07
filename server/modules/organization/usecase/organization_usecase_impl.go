package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/organization/repository"
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/jinzhu/gorm"
)

// OrganizationUsecaseImpl struct
type OrganizationUsecaseImpl struct {
	organizationRepositoryRead  repository.OrganizationRepository
	organizationRepositoryWrite repository.OrganizationRepository
}

// NewOrganizationUsecaseImpl function
func NewOrganizationUsecaseImpl(OrganizationRepositoryRead, OrganizationRepositoryWrite repository.OrganizationRepository) *OrganizationUsecaseImpl {
	return &OrganizationUsecaseImpl{
		organizationRepositoryRead:  OrganizationRepositoryRead,
		organizationRepositoryWrite: OrganizationRepositoryWrite,
	}
}

// CreateOrganization function
func (u *OrganizationUsecaseImpl) CreateOrganization(organization *domain.Organization) shared.Output {
	organizationOutput := u.organizationRepositoryRead.FindByID(organization.ID)
	if organizationOutput.Err != nil && organizationOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output{Err: organizationOutput.Err}
	}

	if organizationOutput.Result != nil {
		organizationExist := organizationOutput.Result.(*domain.Organization)

		if organizationExist != nil {
			return shared.Output{Err: fmt.Errorf("organization with id: %d already exist", organization.ID)}
		}
	}

	err := organization.Validate()

	if err != nil {
		return shared.Output{Err: err}
	}

	organizationSaveOutput := u.organizationRepositoryWrite.Save(organization)
	if organizationSaveOutput.Err != nil {
		return shared.Output{Err: organizationSaveOutput.Err}
	}

	organizationSaved := organizationSaveOutput.Result.(*domain.Organization)

	return shared.Output{Result: organizationSaved}
}

// RemoveOrganization function
func (u *OrganizationUsecaseImpl) RemoveOrganization(id uint) shared.Output {
	organizationResult := u.organizationRepositoryRead.FindByID(id)
	if organizationResult.Err != nil {
		return shared.Output{Err: organizationResult.Err}
	}

	organization := organizationResult.Result.(*domain.Organization)

	organizationDeleteOutput := u.organizationRepositoryWrite.Delete(organization)
	if organizationDeleteOutput.Err != nil {
		return shared.Output{Err: organizationDeleteOutput.Err}
	}

	organizationDeleted := organizationDeleteOutput.Result.(*domain.Organization)
	return shared.Output{Result: organizationDeleted}
}

// GetOrganization function
func (u *OrganizationUsecaseImpl) GetOrganization(id uint) shared.Output {
	organizationResult := u.organizationRepositoryRead.FindByID(id)
	if organizationResult.Err != nil {
		return shared.Output{Err: organizationResult.Err}
	}

	organization := organizationResult.Result.(*domain.Organization)

	return shared.Output{Result: organization}
}

// GetAllOrganization function
func (u *OrganizationUsecaseImpl) GetAllOrganization(params *shared.Parameters) shared.Output {
	params.Page = 1

	if len(params.StrPage) > 0 {
		page, err := strconv.Atoi(params.StrPage)
		if err != nil {
			return shared.Output{Err: shared.NewErrorAllowNumericOnly("page")}
		}

		params.Page = page
	}

	params.Limit = 10
	if len(params.StrLimit) > 0 {
		limit, err := strconv.Atoi(params.StrLimit)
		if err != nil {
			return shared.Output{Err: shared.NewErrorAllowNumericOnly("limit")}
		}

		params.Limit = limit
	}

	params.Offset = (params.Page - 1) * params.Limit

	if len(params.OrderBy) > 0 {
		if !shared.StringInSlice(params.OrderBy, shared.AllowedSortFields) {
			return shared.Output{Err: fmt.Errorf(shared.ErrorParameterInvalid, "order by")}
		}
	} else {
		params.OrderBy = "name"
	}

	if len(params.Sort) > 0 {
		if !shared.StringInSlice(params.Sort, []string{"asc", "desc"}) {
			return shared.Output{Err: fmt.Errorf("parameter %s allow input asc and desc only", "sort")}

		}
	} else {
		params.Sort = "asc"
	}

	params.OrderBy = fmt.Sprintf(`"%s" %s`, strings.ToUpper(params.OrderBy), params.Sort)

	organizationResult := u.organizationRepositoryRead.FindAll(params)
	if organizationResult.Err != nil {
		return shared.Output{Err: organizationResult.Err}
	}

	organizations := organizationResult.Result.(domain.Organizations)

	return shared.Output{Result: organizations}
}

// GetTotalOrganization function
func (u *OrganizationUsecaseImpl) GetTotalOrganization(params *shared.Parameters) shared.Output {
	organizationResult := u.organizationRepositoryRead.Count(params)
	if organizationResult.Err != nil {
		return shared.Output{Err: organizationResult.Err}
	}

	totalOrganization := organizationResult.Result.(int)

	return shared.Output{Result: totalOrganization}
}
