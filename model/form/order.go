package form

const (
	OrderTypeNotPay = iota + 1
	OrderTypePayed
	OrderTypeDone
)

type Order struct {
	UserID    uint    `json:"user_id,omitempty" bson:"user_id"`
	ProductID uint    `json:"product_id,omitempty" bson:"product_id"`
	BossID    uint    `json:"boss_id,omitempty" bson:"boss_id"`
	AddressID uint    `json:"address_id,omitempty" bson:"address_id"`
	Num       int     `json:"num,omitempty" bson:"num"`             // 数量
	OrderNum  uint64  `json:"order_num,omitempty" bson:"order_num"` // 订单号
	Type      uint    `json:"type,omitempty" bson:"type"`           // 1 未支付  2 已支付  3 已完成
	Money     float64 `json:"money,omitempty" bson:"money"`
	Id        uint    `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt int64   `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64   `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt  int64   `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
}
