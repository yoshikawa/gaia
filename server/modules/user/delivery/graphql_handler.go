package delivery

import (
	"context"
	"time"

	"github.com/Pluslab/gaia/server/modules/user/domain"
	"github.com/Pluslab/gaia/server/modules/user/graphql/schema"
	"github.com/Pluslab/gaia/server/modules/user/usecase"
	"github.com/graph-gophers/graphql-go"
)

// GraphQLUserHandler struct
// Handler means Resolver
type GraphQLUserHandler struct {
	UserUsecase usecase.UserUsecase
}

// UserInputArgs input
type UserInputArgs struct {
	User schema.UserSchemaInput
}

// DeleteUserArgs input
type DeleteUserArgs struct {
	ID graphql.ID
}

// UserQueryArgs args
type UserQueryArgs struct {
	ID graphql.ID
}

// UsersArgs struct
type UsersArgs struct {
	Query   *string
	Limit   *float64
	Page    *float64
	OrderBy *string
	Sort    *string
}

// UsersQueryArgs args
type UsersQueryArgs struct {
	UsersArgs *UsersArgs
}

// CreateUser mutation
func (r *GraphQLUserHandler) CreateUser(ctx context.Context, args *UserInputArgs) (*schema.UserSchema, error) {
	var user domain.User
	user.Name = args.User.Name
	user.OrganizationID = args.User.OrganizationID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	output := r.UserUsecase.CreateUser(&user)
	if output.Err != nil {
		return nil, output.Err
	}

	userSaved := output.Result.(*domain.User)

	return &schema.UserSchema{User: userSaved}, nil
}
