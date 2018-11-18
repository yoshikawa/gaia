package infrastructure

import (
	"github.com/Pluslab/fieldsensing/server/interfaces/controller"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	organizationController := controller.NewOrganizationController(NewSQLHandler())

	router.POST("/organizations", func(c *gin.Context) { organizationController.Create(c) })
	router.GET("/organizations", func(c *gin.Context) { organizationController.Index(c) })
	router.GET("/organizations/:id", func(c *gin.Context) { organizationController.Show(c) })

	Router = router
}
