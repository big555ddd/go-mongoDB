package model

import (
	"time"
)

type CreateUpdateUnixTimestamp struct {
	CreateUnixTimestamp `bson:",inline"`
	UpdateUnixTimestamp `bson:",inline"`
}

type CreateUnixTimestamp struct {
	CreatedAt int64 `json:"created_at" bson:"created_at"` // ลบ bun tags
}

type UpdateUnixTimestamp struct {
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"` // ลบ bun tags
}

func (t *CreateUnixTimestamp) SetCreated(ts int64) {
	t.CreatedAt = ts
}

func (t *CreateUnixTimestamp) SetCreatedNow() {
	t.SetCreated(time.Now().Unix())
}

func (t *UpdateUnixTimestamp) SetUpdate(ts int64) {
	t.UpdatedAt = ts
}

func (t *UpdateUnixTimestamp) SetUpdateNow() {
	t.SetUpdate(time.Now().Unix())
}

// Unix Milli
type CreateUpdateMilliTimestamp struct {
	CreateMilliTimestamp
	UpdateMilliTimestamp
}

type CreateMilliTimestamp struct {
	CreatedAt int64 `json:"created_at" bson:"created_at"` // ลบ bun tags
}

type UpdateMilliTimestamp struct {
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"` // ลบ bun tags
}

func (t *CreateMilliTimestamp) SetCreated(ts int64) {
	t.CreatedAt = ts
}

func (t *CreateMilliTimestamp) SetCreatedNow() {
	t.SetCreated(time.Now().UnixMilli())
}

func (t *UpdateMilliTimestamp) SetUpdate(ts int64) {
	t.UpdatedAt = ts
}

func (t *UpdateMilliTimestamp) SetUpdateNow() {
	t.SetUpdate(time.Now().UnixMilli())
}
