package routes

import (
	"app/app/controller"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	// Get the *bun.DB instance from config
	ctl := controller.New() // Pass the *bun.DB to the controller
	// md := middleware.AuthMiddleware()
	// log := middleware.NewLogResponse()
	user := router.Group("")
	{
		user.GET("/list", ctl.UserCtl.List)
		user.GET("/:id", ctl.UserCtl.Get)
		user.PATCH("/:id", ctl.UserCtl.Update)
		user.DELETE("/:id", ctl.UserCtl.Delete)
		user.POST("/create", ctl.UserCtl.Create)
	}
}
