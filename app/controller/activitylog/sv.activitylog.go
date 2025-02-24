package activitylog

import (
	"app/app/model"
	"app/internal/logger"
	"context"
)

// Ensure to import the model package

func (s *Service) Create(ctx context.Context, req model.ActivityLog) (*model.ActivityLog, error) {
	if _, err := s.db.Collection("activity_logs").InsertOne(ctx, &req); err != nil {
		logger.Err(err.Error())
		return nil, err
	}
	return &req, nil
}
