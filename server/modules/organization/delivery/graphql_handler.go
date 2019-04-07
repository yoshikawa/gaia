package delivery

import (
	"context"
	"math"
	"strconv"
	"time"

	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/organization/graphql/schema"
	"github.com/Pluslab/gaia/server/modules/organization/usecase"
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/graph-gophers/graphql-go"
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

// DeleteOrganizationArgs input
type DeleteOrganizationArgs struct {
	ID graphql.ID
}

// OrganizationQueryArgs args
type OrganizationQueryArgs struct {
	ID graphql.ID
}

// OrganizationsArgs struct
type OrganizationsArgs struct {
	Query   *string
	Limit   *float64
	Page    *float64
	OrderBy *string
	Sort    *string
}

// OrganizationsQueryArgs args
type OrganizationsQueryArgs struct {
	OrganizationsArgs *OrganizationsArgs
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

// DeleteOrganization mutation
func (r *GraphQLOrganizationHandler) DeleteOrganization(ctx context.Context, args *DeleteOrganizationArgs) (*schema.OrganizationSchema, error) {
	stringID := string(args.ID)
	argsID, err := (strconv.ParseUint(stringID, 10, 32))

	if err != nil {
		return nil, err
	}

	output := r.OrganizationUsecase.RemoveOrganization(uint(argsID))

	if output.Err != nil {
		return nil, output.Err
	}

	organization := output.Result.(*domain.Organization)

	return &schema.OrganizationSchema{Organization: organization}, nil
}

// Organization Query function
func (r *GraphQLOrganizationHandler) Organization(ctx context.Context, args *OrganizationQueryArgs) (*schema.OrganizationSchema, error) {
	stringID := string(args.ID)
	argsID, err := (strconv.ParseUint(stringID, 10, 32))

	if err != nil {
		return nil, err
	}

	output := r.OrganizationUsecase.GetOrganization(uint(argsID))

	if output.Err != nil {
		return nil, output.Err
	}

	organization := output.Result.(*domain.Organization)

	return &schema.OrganizationSchema{Organization: organization}, nil

}

// Organizations Query function
func (r *GraphQLOrganizationHandler) Organizations(ctx context.Context, args *OrganizationsQueryArgs) (*schema.OrganizationListResolver, error) {
	var (
		params                shared.Parameters
		organizationsResolver []*schema.OrganizationSchema
		meta                  schema.MetaResolver
		result                schema.OrganizationListResolver
	)
	if args.OrganizationsArgs.Limit != nil {
		limitStr := strconv.Itoa(int(*args.OrganizationsArgs.Limit))
		params.StrLimit = limitStr
	}

	if args.OrganizationsArgs.Page != nil {
		pageStr := strconv.Itoa(int(*args.OrganizationsArgs.Page))
		params.StrPage = pageStr
	}

	if args.OrganizationsArgs.OrderBy != nil {
		params.OrderBy = *args.OrganizationsArgs.OrderBy
	}

	if args.OrganizationsArgs.Sort != nil {
		params.Sort = *args.OrganizationsArgs.Sort
	}

	organizationsOutput := r.OrganizationUsecase.GetAllOrganization(&params)

	if organizationsOutput.Err != nil {
		return nil, organizationsOutput.Err
	}

	organizations := organizationsOutput.Result.(domain.Organizations)

	if len(organizations) <= 0 {
		limitInt32 := int32(params.Limit)
		pageInt32 := int32(params.Page)
		totalInt32 := int32(0)
		totalPageInt32 := int32(0)

		meta.LimitField = &limitInt32
		meta.PageField = &pageInt32
		meta.TotalRecordsField = &totalInt32
		meta.TotalPagesField = &totalPageInt32

		result.OrganizationsField = organizationsResolver
		result.MetaField = &meta

		return &result, nil
	}

	for _, organization := range organizations {
		organizationsResolver = append(organizationsResolver, &schema.OrganizationSchema{Organization: organization})
	}

	countOutput := r.OrganizationUsecase.GetTotalOrganization(&params)

	if countOutput.Err != nil {
		return nil, countOutput.Err
	}

	total := countOutput.Result.(int)

	totalPage := int(math.Ceil(float64(total) / float64(params.Limit)))

	limitInt32 := int32(params.Limit)
	pageInt32 := int32(params.Page)
	totalInt32 := int32(total)
	totalPageInt32 := int32(totalPage)

	meta.LimitField = &limitInt32
	meta.PageField = &pageInt32
	meta.TotalRecordsField = &totalInt32
	meta.TotalPagesField = &totalPageInt32

	result.OrganizationsField = organizationsResolver
	result.MetaField = &meta

	return &result, nil
}
