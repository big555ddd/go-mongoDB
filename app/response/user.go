package response

type UserResponse struct {
	ID          string `bson:"_id,omitempty" json:"id"` // เปลี่ยนเป็น _id สำหรับ MongoDB
	Username    string `bson:"username" json:"username"`
	Email       string `bson:"email" json:"email"`
	FirstName   string `bson:"first_name" json:"first_name"`
	LastName    string `bson:"last_name" json:"last_name"`
	DisplayName string `bson:"display_name" json:"display_name"`
	RoleID      int64  `bson:"role_id" json:"role_id"`
	Status      string `bson:"status" json:"status"`
}
