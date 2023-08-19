package form

type Cart struct {
	Id        uint  `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt  int64 `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
	UserID    uint  `json:"user_id,omitempty" bson:"user_id"`
	ProductID uint  `json:"product_id,omitempty" bson:"product_id"`
	BossID    uint  `json:"boss_id,omitempty" bson:"boss_id"`
	Num       uint  `json:"num,omitempty" bson:"num"`
	MaxNum    uint  `json:"max_num,omitempty" bson:"max_num"`
	Check     bool  `json:"check,omitempty" bson:"check"`
}
