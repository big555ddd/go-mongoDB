package request

type LoginUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type CreateUser struct {
	Username    string `json:"username" bson:"username"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password" bson:"password"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	DisplayName string `json:"display_name" bson:"display_name"`
}

type UpdateUser struct {
	Password    string `json:"password" bson:"password"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	DisplayName string `json:"display_name" bson:"display_name"`
}

type ResetPassword struct {
	Email string `json:"email" bson:"email"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password" bson:"old_password"`
	NewPassword string `json:"new_password" bson:"new_password"`
}
