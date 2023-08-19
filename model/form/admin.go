package form

type Admin struct {
	Id             uint   `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt       int64  `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
	UserName       string `json:"user_name,omitempty" bson:"user_name"`
	PasswordDigest string `json:"password_digest,omitempty" bson:"password_digest"`
	Avatar         string `json:"avatar,omitempty" bson:"avatar"`
}
