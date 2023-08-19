package form

type Product struct {
	Id            uint   `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt      int64  `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
	Name          string `json:"name,omitempty" bson:"name"`
	CategoryID    uint   `json:"category_id,omitempty" bson:"category_id"`
	Title         string `json:"title,omitempty" bson:"title"`
	Info          string `json:"info,omitempty" bson:"info"`
	ImgPath       string `json:"img_path,omitempty" bson:"img_path"`
	Price         string `json:"price,omitempty" bson:"price"`
	DiscountPrice string `json:"discount_price,omitempty" bson:"discount_price"`
	OnSale        bool   `json:"on_sale,omitempty" bson:"on_sale"`
	Num           int    `json:"num,omitempty" bson:"num"`
	BossID        uint   `json:"boss_id,omitempty" bson:"boss_id"`
	BossName      string `json:"boss_name,omitempty" bson:"boss_name"`
	BossAvatar    string `json:"boss_avatar,omitempty" bson:"boss_avatar"`
}
