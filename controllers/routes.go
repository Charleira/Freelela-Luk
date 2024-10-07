package controllers

import (
	drivercontroller "github.com/Charleira/FreelelaLuk/controllers/driver-controller"
	"github.com/gin-gonic/gin"
)

func AddRoutes(app *gin.Engine) {

	// Drivers
	app.POST("/create-users", drivercontroller.CreateUser)

}
