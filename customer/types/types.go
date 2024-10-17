package types

import (
	"time"
)

type CustomerRequestModel struct {
	FirstName    string        `bson:"first_name" json:"first_name" validate:"required,min=2,max=18"`
	LastName     string        `bson:"last_name" json:"last_name" validate:"required,min=2,max=20"`
	Age          int           `bson:"age" json:"age" validate:"required,min=18,max=75"`
	Email        string        `bson:"email" json:"email" validate:"required,email"`
	Username     string        `bson:"username" json:"username" validate:"required"`
	Password     string        `bson:"password" json:"password" validate:"required"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
	Addresses    []Address     `bson:"addresses" json:"addresses"`
	PhoneNumbers []PhoneNumber `bson:"phone_numbers" json:"phone_numbers"`
}

type Address struct {
	ID         string `bson:"_id,omitempty" json:"id"`
	CustomerId string `bson:"customer_id" json:"customer_id"`
	Street     string `bson:"street" json:"street" validate:"required"`
	City       string `bson:"city" json:"city" validate:"required"`
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
	FirstName      string            `bson:"first_name" json:"first_name"`
	LastName       string            `bson:"last_name" json:"last_name"`
	Username       string            `bson:"username" json:"username"`
	Password       string            `bson:"password" json:"password"`
	Age            int               `bson:"age" json:"age"`
	Email          string            `bson:"email" json:"email"`
	Phone          string            `bson:"phone" json:"phone"`
	AdditionalInfo map[string]string `bson:"additional_info" json:"additional_info"`
	ContactOption  []string          `bson:"contact_option" json:"contact_option"`
	Token          string            `bson:"token" json:"token"`
	Addresses      []Address         `bson:"addresses" json:"addresses"`
	PhoneNumber    []PhoneNumber     `bson:"phone_numbers" json:"phone_numbers"`
}

type CustomerUpdateModel struct {
	FirstName      string            `bson:"first_name" json:"first_name"`
	LastName       string            `bson:"last_name" json:"last_name"`
	Age            int               `bson:"age" json:"age"`
	Phone          string            `bson:"phone" json:"phone"`
	Address        string            `bson:"address" json:"address"`
	City           string            `bson:"city" json:"city"`
	State          string            `bson:"state" json:"state"`
	ZipCode        string            `bson:"zip_code" json:"zip_code"`
	Username       string            `bson:"username" json:"username" validate:"required"`
	Password       string            `bson:"password" json:"password" validate:"required"`
	AdditionalInfo map[string]string `bson:"additional_info" json:"additional_info"`
	MembershipType string            `bson:"membership_type" json:"membership_type"`
	ContactOption  []string          `bson:"contact_option" json:"contact_option"`
	CreatedAt      time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time         `bson:"updated_at" json:"updated_at"`
}

func ToCustomerResponse(customer *Customer) *CustomerResponseModel {
	return &CustomerResponseModel{
		FirstName:      customer.FirstName,
		LastName:       customer.LastName,
		Username:       customer.Username,
		Password:       customer.Password,
		Age:            customer.Age,
		Email:          customer.Email,
		PhoneNumber:    customer.PhoneNumbers,
		AdditionalInfo: customer.AdditionalInfo,
		ContactOption:  customer.ContactOption,
		Addresses:      customer.Addresses,
	}
}
