package server

import (
	"fmt"
	"net/http"

	"github.com/Pluslab/gaia/server/database"
	cd "github.com/Pluslab/gaia/server/modules/organization/delivery"
	cr "github.com/Pluslab/gaia/server/modules/organization/repository"
	cu "github.com/Pluslab/gaia/server/modules/organization/usecase"
	pd "github.com/Pluslab/gaia/server/modules/user/delivery"
	pr "github.com/Pluslab/gaia/server/modules/user/repository"
	pu "github.com/Pluslab/gaia/server/modules/user/usecase"
	"github.com/Pluslab/gaia/server/schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
)

// embedding all graphql resolver/ handler anonymously
type graphqlHandlers struct {
	pd.GraphQLUserHandler
	cd.GraphQLOrganizationHandler
}

// EchoServer struct
type EchoServer struct {
	port           int
	graphQLHandler *relay.Handler
}

// NewEchoServer echo server constructor
func NewEchoServer(port int) (*EchoServer, error) {
	db, err := database.GetGormConn()
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	// load graphql schema file, and convert to string
	graphqlSchema, err := schema.LoadGraphQLSchema()
	if err != nil {
		return nil, err
	}

	// initial repository
	userRepository := pr.NewUserRepositoryGorm(db)
	organizationRepository := cr.NewOrganizationRepositoryGorm(db)

	// initial usecase
	userUsecase := pu.NewUserUsecaseImpl(userRepository, userRepository, organizationRepository)
	organizationUsecase := cu.NewOrganizationUsecaseImpl(organizationRepository, organizationRepository)

	// initial graphql handler/ resolver
	userGraphQLHandler := pd.GraphQLUserHandler{UserUsecase: userUsecase}
	organizationGraphQLHandler := cd.GraphQLOrganizationHandler{OrganizationUsecase: organizationUsecase}

	// create graphql resolver
	var graphqlResolver graphqlHandlers

	graphqlResolver.GraphQLUserHandler = userGraphQLHandler
	graphqlResolver.GraphQLOrganizationHandler = organizationGraphQLHandler

	// parse grapqhql schema to code
	gqlSchema := graphql.MustParseSchema(graphqlSchema, &graphqlResolver)

	graphQLHandler := &relay.Handler{Schema: gqlSchema}

	return &EchoServer{
		port:           port,
		graphQLHandler: graphQLHandler,
	}, nil
}

// Run function
func (s *EchoServer) Run() {
	e := echo.New()
	e.Use(echoMiddleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Up and running !!")
	})

	// secure graphql route with Basic Auth
	e.POST("/graphql", echo.WrapHandler(s.graphQLHandler))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.port)))

}
