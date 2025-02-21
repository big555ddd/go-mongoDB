package activitylog

import "app/app/provider/database"

type Controller struct {
	Name    string
	Service *Service
}

func NewController(db *database.MongoDB) *Controller {
	return &Controller{
		Name:    `activity-log-ctl`,
		Service: NewService(db),
	}
}

type Service struct {
	db *database.MongoDB
}

func NewService(db *database.MongoDB) *Service {
	return &Service{
		db: db,
	}
}
