package controllers

import (
	drivercontroller "github.com/Charleira/FreelelaLuk/controllers/driver-controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(app *gin.Engine) {

	// Drivers
	app.POST("/create-users", drivercontroller.CreateUser)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
