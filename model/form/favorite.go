package form

type Favorite struct {
	Id        uint    `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt int64   `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64   `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt  int64   `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
	User      User    `json:"user" bson:"user"`
	UserID    uint    `json:"user_id,omitempty" bson:"user_id"`
	Product   Product `json:"product,omitempty" bson:"product"`
	ProductID uint    `json:"product_id,omitempty" bson:"product_id"`
	Boss      User    `json:"boss" bson:"boss"`
	BossID    uint    `json:"boss_id,omitempty" bson:"boss_id"`
}
