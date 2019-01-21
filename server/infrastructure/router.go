package infrastructure

import (
	"github.com/Pluslab/gaia/server/interfaces/controller"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Router is gin engine
var Router *gin.Engine

func init() {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	// SessionNameは任意
	router.Use(sessions.Sessions("SessionName", store))

	sessionController := controller.NewSessionController(NewSQLHandler())
	userController := controller.NewUserController(NewSQLHandler())
	organizationController := controller.NewOrganizationController(NewSQLHandler())
	vendorController := controller.NewVendorController(NewSQLHandler())
	fieldController := controller.NewFieldController(NewSQLHandler())
	gatewayDeviceController := controller.NewGatewayDeviceController(NewSQLHandler())
	gatewayDeviceStatusController := controller.NewGatewayDeviceStatusController(NewSQLHandler())
	individualController := controller.NewIndividualController(NewSQLHandler())
	observablePropertyController := controller.NewObservablePropertyController(NewSQLHandler())
	observableDatumController := controller.NewObservationDatumController(NewSQLHandler())
	observationPositionController := controller.NewObservationPositionController(NewSQLHandler())
	plantCategoryController := controller.NewPlantCategoryController(NewSQLHandler())
	plantController := controller.NewPlantController(NewSQLHandler())
	sensingDeviceController := controller.NewSensingDeviceController(NewSQLHandler())
	sensingDeviceStatusController := controller.NewSensingDeviceStatusController(NewSQLHandler())
	sensorController := controller.NewSensorController(NewSQLHandler())

	// session awi route
	router.POST("/login", func(c *gin.Context) { sessionController.Login(c) })
	router.GET("/logout", func(c *gin.Context) { sessionController.Logout(c) })
	// user api route
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	// organization api route
	router.POST("/organizations", func(c *gin.Context) { organizationController.Create(c) })
	router.GET("/organizations", func(c *gin.Context) { organizationController.Index(c) })
	router.GET("/organizations/:id", func(c *gin.Context) { organizationController.Show(c) })
	// vendor api route
	router.POST("/vendors", func(c *gin.Context) { vendorController.Create(c) })
	router.GET("/vendors", func(c *gin.Context) { vendorController.Index(c) })
	router.GET("/vendors/:id", func(c *gin.Context) { vendorController.Show(c) })
	// field api route
	router.POST("/fields", func(c *gin.Context) { fieldController.Create(c) })
	router.GET("/fields", func(c *gin.Context) { fieldController.Index(c) })
	router.GET("/fields/:id", func(c *gin.Context) { fieldController.Show(c) })
	// gateway device api route
	router.POST("/gateway-devices", func(c *gin.Context) { gatewayDeviceController.Create(c) })
	router.GET("/gateway-devices", func(c *gin.Context) { gatewayDeviceController.Index(c) })
	router.GET("/gateway-devices/:id", func(c *gin.Context) { gatewayDeviceController.Show(c) })
	// gateway device status api route
	router.POST("/gateway-device-statuses", func(c *gin.Context) { gatewayDeviceStatusController.Create(c) })
	router.GET("/gateway-device-statuses", func(c *gin.Context) { gatewayDeviceStatusController.Index(c) })
	router.GET("/gateway-device-statuses/:id", func(c *gin.Context) { gatewayDeviceStatusController.Show(c) })
	// individual api route
	router.POST("/individuals", func(c *gin.Context) { individualController.Create(c) })
	router.GET("/individuals", func(c *gin.Context) { individualController.Index(c) })
	router.GET("/individuals/:id", func(c *gin.Context) { individualController.Show(c) })
	// observable property api route
	router.POST("/observable-properties", func(c *gin.Context) { observablePropertyController.Create(c) })
	router.GET("/observable-properties", func(c *gin.Context) { observablePropertyController.Index(c) })
	router.GET("/observable-properties/:id", func(c *gin.Context) { observablePropertyController.Show(c) })
	// observable datum api route
	router.POST("/observable-data", func(c *gin.Context) { observableDatumController.Create(c) })
	router.GET("/observable-data", func(c *gin.Context) { observableDatumController.Index(c) })
	router.GET("/observable-data/:id", func(c *gin.Context) { observableDatumController.Show(c) })
	// /observation position api route
	router.POST("/observation-positions", func(c *gin.Context) { observationPositionController.Create(c) })
	router.GET("/observation-positions", func(c *gin.Context) { observationPositionController.Index(c) })
	router.GET("/observation-positions/:id", func(c *gin.Context) { observationPositionController.Show(c) })
	// plant category api route
	router.POST("/plant-categories", func(c *gin.Context) { plantCategoryController.Create(c) })
	router.GET("/plant-categories", func(c *gin.Context) { plantCategoryController.Index(c) })
	router.GET("/plant-categories/:id", func(c *gin.Context) { plantCategoryController.Show(c) })
	// plant api route
	router.POST("/plants", func(c *gin.Context) { plantController.Create(c) })
	router.GET("/plants", func(c *gin.Context) { plantController.Index(c) })
	router.GET("/plants/:id", func(c *gin.Context) { plantController.Show(c) })
	// sensing device api route
	router.POST("/sensing-devices", func(c *gin.Context) { sensingDeviceController.Create(c) })
	router.GET("/sensinsg-devices", func(c *gin.Context) { sensingDeviceController.Index(c) })
	router.GET("/sensinsg-devices/:id", func(c *gin.Context) { sensingDeviceController.Show(c) })
	// sensing device status api route
	router.POST("/sensing-deivce-statuses", func(c *gin.Context) { sensingDeviceStatusController.Create(c) })
	router.GET("/sensing-device-statuses", func(c *gin.Context) { sensingDeviceStatusController.Index(c) })
	router.GET("/sensing-device-statuses/:id", func(c *gin.Context) { sensingDeviceStatusController.Show(c) })
	// sensor api route
	router.POST("/sesors", func(c *gin.Context) { sensorController.Create(c) })
	router.GET("/sensors", func(c *gin.Context) { sensorController.Index(c) })
	router.GET("/sensors/:id", func(c *gin.Context) { sensorController.Show(c) })

	Router = router
}
