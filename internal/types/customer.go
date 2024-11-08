package types

import (
	"time"
)

type CustomerRequestModel struct {
	FirstName     string        `bson:"first_name" json:"first_name" validate:"required,min=2,max=18"`
	LastName      string        `bson:"last_name" json:"last_name" validate:"required,min=2,max=20"`
	Age           int           `bson:"age" json:"age" validate:"required,min=18,max=75"`
	Email         string        `bson:"email" json:"email" validate:"required,email"`
	Username      string        `bson:"username" json:"username" validate:"required"`
	Password      string        `bson:"password" json:"password" validate:"required"`
	Addresses     []Address     `bson:"addresses,omitempty" json:"addresses" validate:"dive,required"`
	PhoneNumbers  []PhoneNumber `bson:"phone_numbers,omitempty" json:"phone_numbers"`
	CreatorUserId string        `bson:"creator_user_id" json:"creator_user_id"`
}

type Address struct {
	Street string `bson:"street" json:"street" validate:"required"`
	City   string `bson:"city" json:"city" validate:"required"`
}

type PhoneNumber struct {
	ID         string `bson:"_id,omitempty" json:"id"`
	CustomerId string `bson:"customer_id" json:"customer_id"`
	Phone      string `bson:"phone" json:"phone" validate:"len=10,numeric"`
}

type QueryParams struct {
	FirstName      string `json:"first_name"`
	AgeGreaterThan string `json:"agt"`
	AgeLessThan    string `json:"alt"`
}

type CustomerResponseModel struct {
	FirstName     string        `bson:"first_name" json:"first_name" validate:"required,min=2,max=18,dive"`
	LastName      string        `bson:"last_name" json:"last_name" validate:"required,min=2,max=20"`
	Age           int           `bson:"age" json:"age" validate:"required,min=18,max=75"`
	Email         string        `bson:"email" json:"email" validate:"required,email"`
	Username      string        `bson:"username" json:"username" validate:"required"`
	Password      string        `bson:"password" json:"password" validate:"required"`
	Addresses     []Address     `bson:"addresses,omitempty" json:"addresses" validate:"dive,required"`
	PhoneNumbers  []PhoneNumber `bson:"phone_numbers,omitempty" json:"phone_numbers"`
	CreatorUserId string        `bson:"creator_user_id" json:"creator_user_id"`
	CreatedAt     time.Time     `bson:"created_at" json:"created_at"`
}

type CustomerUpdateModel struct {
	FirstName string    `bson:"first_name" json:"first_name" validate:"required,min=2,max=18"`
	LastName  string    `bson:"last_name" json:"last_name" validate:"required,min=2,max=20"`
	Age       int       `bson:"age" json:"age" validate:"required,min=18,max=75"`
	Phone     string    `bson:"phone" json:"phone"`
	Address   string    `bson:"address" json:"address"`
	City      string    `bson:"city" json:"city"`
	State     string    `bson:"state" json:"state"`
	Username  string    `bson:"username" json:"username" validate:"required"`
	Password  string    `bson:"password" json:"password" validate:"required"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func ToCustomerResponse(customer *Customer) *CustomerResponseModel {
	return &CustomerResponseModel{
		FirstName:    customer.FirstName,
		LastName:     customer.LastName,
		Age:          customer.Age,
		Email:        customer.Email,
		PhoneNumbers: customer.PhoneNumbers,
		Addresses:    customer.Addresses,
		CreatedAt:    customer.CreatedAt,
	}
}
