package types

import (
	"time"
)

type OrderRequestModel struct {
	CreatorUserId string `bson:"creator_user_id" json:"creator_user_id"`
	OrderTotal    int    `bson:"order_total" json:"order_total"`
	OrderName     string `bson:"order_name" json:"order_name"`
	PaymentMethod string `bson:"payment_method" json:"payment_method"`
}

type OrderResponseModel struct {
	CustomerId    string    `bson:"customer_id" json:"customer_id"`
	OrderName     string    `bson:"order_name" json:"order_name"`
	OrderTotal    int       `bson:"order_total" json:"order_total"`
	CreatorUserId string    `bson:"creator_user_id" json:"creator_user_id"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
}
type OrderUpdateModel struct {
	OrderName     string `bson:"order_name" json:"order_name"`
	OrderTotal    int    `bson:"order_total" json:"order_total"`
	OrderStatus   string `bson:"order_status" json:"order_status"`
	PaymentMethod string `bson:"payment_method" json:"payment_method"`
}

type CustomerResponse struct {
	UserName string `bson:"username" json:"username"`
	Name     string `bson:"first_name" json:"first_name"`
	LastName string `bson:"last_name" json:"last_name"`
	Address  string `bson:"address" json:"address"`
}

func ToOrderResponse(order *Order) *OrderResponseModel {
	return &OrderResponseModel{
		OrderName:     order.OrderName,
		OrderTotal:    order.OrderTotal,
		CreatorUserId: order.CreatorUserId,
		CreatedAt:     order.CreatedAt,
	}
}
