package controllers

import (
	drivercontroller "github.com/Charleira/FreelelaLuk/controllers/driver-controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(app *gin.Engine) {
	// Swagger documentation
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth routes
	app.POST("/auth/register", drivercontroller.RegisterUser)
	app.POST("/auth/login", drivercontroller.LoginUser)

	// Store routes
	app.POST("/store/purchase", drivercontroller.RecordPurchase)
	app.GET("/store/items", drivercontroller.GetStoreItems)

	// Competition routes
	app.POST("/competitions/create", drivercontroller.CreateCompetition)
	app.POST("/competitions/add-team", drivercontroller.AddTeamToCompetition)
	app.GET("/competitions", drivercontroller.ListCompetitions)

	// Admin routes
	app.GET("/admin/competitions", drivercontroller.ListAllCompetitions)
	app.GET("/admin/sales", drivercontroller.ListAllSales)
}
