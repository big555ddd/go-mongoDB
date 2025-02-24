package model

import "github.com/uptrace/bun"

type RolePermission struct {
	bun.BaseModel `bun:"table:role_permissions"`

	RoleID       int64 `bson:"role_id" json:"role_id"`             // FK ใช้ชื่อปกติ
	PermissionID int64 `bson:"permission_id" json:"permission_id"` // FK ใช้ชื่อปกติ

	CreateUpdateUnixTimestamp `bson:",inline"`
}
