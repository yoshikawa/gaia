package schema

import (
	"time"

	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/graph-gophers/graphql-go"
)

// OrganizationSchemaInput struct
type OrganizationSchemaInput struct {
	Name      string
	CreatedAt *time.Time
	UpdateAt  *time.Time
}

// OrganizationSchema resolver
type OrganizationSchema struct {
	Organization *domain.Organization
}

// ID function
func (organization *OrganizationSchema) ID() graphql.ID {
	return graphql.ID(organization.Organization.ID)
}

// Name function
func (organization *OrganizationSchema) Name() string {
	return organization.Organization.Name
}

// CreatedAt function
func (organization *OrganizationSchema) CreatedAt() *graphql.Time {
	return &graphql.Time{Time: organization.Organization.CreatedAt}
}

// UpdatedAt function
func (organization *OrganizationSchema) UpdatedAt() *graphql.Time {
	return &graphql.Time{Time: organization.Organization.UpdatedAt}
}
