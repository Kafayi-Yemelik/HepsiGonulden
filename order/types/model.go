package types

import "time"

type Order struct {
	Id            string `bson:"_id" json:"id"`
	OrderName     string `bson:"order_name,omitempty" json:"order_name"`
	CreatorUserId string `bson:"creator_user_id" json:"creator_user_id"`
	PaymentMethod string `bson:"payment_method,omitempty" json:"payment_method"`
	OrderTotal    int    `bson:"order_total" json:"order_total"`
	OrderStatus   string `bson:"order_status" json:"order_status"`
	// TODO(Gonul): Status property'sini ekle
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
