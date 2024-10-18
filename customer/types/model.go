package types

import "time"

type Customer struct {
	Id             string            `bson:"_id" json:"id"`
	FirstName      string            `bson:"first_name" json:"first_name" validate:"required"`
	LastName       string            `bson:"last_name" json:"last_name" validate:"required"`
	Age            int               `bson:"age" json:"age"`
	Email          string            `bson:"email" json:"email"`
	AdditionalInfo map[string]string `bson:"additional_info" json:"additional_info"`
	MembershipType string            `bson:"membership_type" json:"membership_type"`
	ContactOption  []string          `bson:"contact_option" json:"contact_option"`
	Username       string            `bson:"username" json:"username" validate:"required"`
	Password       string            `bson:"password" json:"password" validate:"required"`
	CreatedAt      time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time         `bson:"updated_at,omitempty" json:"updated_at"`
	Addresses      []Address         `bson:"addresses,omitempty" json:"addresses"`
	PhoneNumbers   []PhoneNumber     `bson:"phone_numbers,omitempty" json:"phone_numbers"`
}
