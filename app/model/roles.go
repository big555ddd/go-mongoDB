package model

import "github.com/uptrace/bun"

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	ID          int64  `bson:",pk,autoincrement" json:"id"` // ใช้ ID สำหรับ Primary Key
	RoleName    string `bson:"role_name" json:"role_name"`
	Description string `bson:"description" json:"description"`

	CreateUpdateUnixTimestamp `bson:",inline"`
}
