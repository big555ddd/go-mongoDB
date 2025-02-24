package model

import "github.com/uptrace/bun"

type Permission struct {
	bun.BaseModel `bun:"table:permissions"`

	ID             int64  `bson:",pk,autoincrement" json:"id"`
	PermissionName string `bson:"permission_name" json:"permission_name"`
	Description    string `bson:"description" json:"description"`

	CreateUpdateUnixTimestamp `bson:",inline"`
}
