package infrastructure

import (
	"github.com/Pluslab/fieldsensing/server/interfaces/controller"
	"github.com/gin-gonic/gin"
)

// Router is gin engine
var Router *gin.Engine

func init() {
	router := gin.Default()

	organizationController := controller.NewOrganizationController(NewSQLHandler())
	userController := controller.NewUserController(NewSQLHandler())

	// organization api route
	router.POST("/organizations", func(c *gin.Context) { organizationController.Create(c) })
	router.GET("/organizations", func(c *gin.Context) { organizationController.Index(c) })
	router.GET("/organizations/:id", func(c *gin.Context) { organizationController.Show(c) })
	// user api route
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
