package form

type Notice struct {
	Id        uint   `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt  int64  `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
	Text      string `json:"text" bson:"text"`
}
