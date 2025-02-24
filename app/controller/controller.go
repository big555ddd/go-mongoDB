package controller

import (
	"app/app/controller/user"
	"app/config"
)

type Controller struct {
	UserCtl *user.Controller

	// Other controllers...
}

func New() *Controller {
	db := config.GetDB()
	return &Controller{

		UserCtl: user.NewController(db),

		// Other controllers...
	}
}
