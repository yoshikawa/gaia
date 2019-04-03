package delivery

import (
	"context"
	"time"

	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/organization/graphql/schema"
	"github.com/Pluslab/gaia/server/modules/organization/usecase"
)

// GraphQLOrganizationHandler struct
// Handler means Resolver
type GraphQLOrganizationHandler struct {
	OrganizationUsecase usecase.OrganizationUsecase
}

// OrganizationInputArgs input
type OrganizationInputArgs struct {
	Organization schema.OrganizationSchemaInput
}

// CreateOrganization mutation
func (r *GraphQLOrganizationHandler) CreateOrganization(ctx context.Context, args *OrganizationInputArgs) (*schema.OrganizationSchema, error) {
	var organization domain.Organization
	organization.Name = args.Organization.Name
	organization.CreatedAt = time.Now()
	organization.UpdatedAt = time.Now()

	output := r.OrganizationUsecase.CreateOrganization(&organization)
	if output.Err != nil {
		return nil, output.Err
	}

	organizationSaved := output.Result.(*domain.Organization)

	return &schema.OrganizationSchema{Organization: organizationSaved}, nil
}
