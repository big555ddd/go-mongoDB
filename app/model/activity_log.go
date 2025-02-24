package model

type ActivityLog struct {
	ID         string      `json:"id" bson:"_id,omitempty" form:"id"`
	Section    string      `json:"section" bson:"section" form:"section"`
	EventType  string      `json:"event_type" bson:"event_type" form:"event_type"`
	StatusCode int         `json:"status_code" bson:"status_code" form:"status_code"`
	Parameters interface{} `json:"parameters" bson:"parameters" form:"parameters"`
	Responses  interface{} `json:"responses" bson:"responses" form:"responses"`
	Query      interface{} `json:"query" bson:"query" form:"query"`
	IpAddress  string      `json:"ip_address" bson:"ip_address" form:"ip_address"`
	UserAgent  string      `json:"user_agent" bson:"user_agent" form:"user_agent"`
	CreatedBy  string      `json:"created_by" bson:"created_by" form:"created_by"`
	CreatedAt  int64       `json:"created_at" bson:"created_at" form:"created_at"`
}
