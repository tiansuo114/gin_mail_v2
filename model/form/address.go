package form

type Address struct {
	Id        uint   `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt  int64  `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
	UserID    uint   `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Phone     string `json:"phone,omitempty" bson:"phone,omitempty"`
	Address   string `json:"address,omitempty" bson:"address,omitempty"`
}
