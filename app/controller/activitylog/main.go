package activitylog

import "go.mongodb.org/mongo-driver/mongo"

type Controller struct {
	Name    string
	Service *Service
}

func NewController(db *mongo.Database) *Controller {
	return &Controller{
		Name:    `activity-log-ctl`,
		Service: NewService(db),
	}
}

type Service struct {
	db *mongo.Database
}

func NewService(db *mongo.Database) *Service {
	return &Service{
		db: db,
	}
}
