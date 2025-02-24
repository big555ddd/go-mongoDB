package user

import (
	"app/app/enum"
	"app/app/model"
	"app/app/request"
	"app/app/response"
	utils "app/app/util"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Create(ctx context.Context, req request.CreateUser) (*model.User, error) {
	// Check if username exists
	count, err := s.db.Collection("users").CountDocuments(ctx, bson.M{"username": req.Username})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("username already exists")
	}

	// Check if email exists
	count, err = s.db.Collection("users").CountDocuments(ctx, bson.M{"email": req.Email})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &model.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    string(hashedPassword),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DisplayName: req.DisplayName,
		RoleID:      1,
		Status:      enum.STATUS_ACTIVE,
	}
	user.SetCreatedNow()
	user.SetUpdateNow()

	// Insert to MongoDB
	_, err = s.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) List(ctx context.Context, limit, page int, search string, roleID string, status string, planType string) ([]response.UserResponse, int, error) {
	var offset int64
	if page > 1 {
		offset = int64((page - 1) * limit)
	}

	filter := bson.M{}
	if search != "" {
		filter["$or"] = []bson.M{
			{"first_name": bson.M{"$regex": search, "$options": "i"}},
			{"display_name": bson.M{"$regex": search, "$options": "i"}},
		}
	}
	if roleID != "" {
		filter["role_id"] = roleID
	}
	if status != "" {
		filter["status"] = status
	}

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(offset).
		SetSort(bson.M{"created_at": 1})

	cursor, err := s.db.Collection("users").Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var res []response.UserResponse
	if err = cursor.All(ctx, &res); err != nil {
		return nil, 0, err
	}

	// Get total count
	total, err := s.db.Collection("users").CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return res, int(total), nil
}

func (s *Service) Get(ctx context.Context, id primitive.ObjectID) (*response.UserResponse, error) {
	res := new(response.UserResponse)
	err := s.db.Collection("users").FindOne(ctx, bson.M{"_id": id}).Decode(res)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Update(ctx context.Context, req request.UpdateUser, id primitive.ObjectID) (*model.User, error) {
	var user model.User
	err := s.db.Collection("users").FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	user = model.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DisplayName: req.DisplayName,
	}

	user.SetCreated(user.CreatedAt)
	user.SetUpdateNow()

	_, err = s.db.Collection("users").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := s.db.Collection("users").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
