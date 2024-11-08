package types

import "time"

type Customer struct {
	Id            string        `bson:"_id" json:"id"`
	FirstName     string        `bson:"first_name" json:"first_name" validate:"required"`
	LastName      string        `bson:"last_name" json:"last_name" validate:"required"`
	Age           int           `bson:"age" json:"age"`
	Email         string        `bson:"email" json:"email"`
	Username      string        `bson:"username" json:"username" validate:"required"`
	Password      string        `bson:"password" json:"password" validate:"required"`
	CreatedAt     time.Time     `bson:"created_at" json:"created_at"`
	CreatorUserId string        `bson:"creator_user_id" json:"creator_user_id" validate:"required"`
	UpdatedAt     time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
	Addresses     []Address     `bson:"addresses,omitempty" json:"addresses"`
	PhoneNumbers  []PhoneNumber `bson:"phone_numbers,omitempty" json:"phone_numbers"`
}
type Order struct {
	Id            string    `bson:"_id" json:"id"`
	OrderName     string    `bson:"order_name,omitempty" json:"order_name"`
	CreatorUserId string    `bson:"creator_user_id" json:"creator_user_id"`
	PaymentMethod string    `bson:"payment_method,omitempty" json:"payment_method"`
	OrderTotal    int       `bson:"order_total" json:"order_total"`
	OrderStatus   string    `bson:"order_status" json:"order_status"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
