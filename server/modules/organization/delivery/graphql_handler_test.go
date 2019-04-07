package delivery

import (
	"context"
	"testing"

	"github.com/Pluslab/gaia/server/modules/organization/graphql/schema"
	usecaseMock "github.com/Pluslab/gaia/server/modules/organization/usecase/mock"
)

func TestGraphQLHandler(t *testing.T) {

	t.Run("should return success test mutation create organization", func(t *testing.T) {
		organizationUsecaseMock := usecaseMock.NewOrganizationUsecaseMock()

		handler := &GraphQLOrganizationHandler{
			OrganizationUsecase: organizationUsecaseMock,
		}

		ctx := context.Background()

		organizationInputArgs := &OrganizationInputArgs{
			Organization: schema.OrganizationSchemaInput{
				Name: "pluslab",
			},
		}

		organizationCreated, err := handler.CreateOrganization(ctx, organizationInputArgs)

		if err != nil {
			t.Error("create organization mutation should return success")
		}

		if organizationCreated == nil {
			t.Error("organization created should not be nil")
		}

	})

}
