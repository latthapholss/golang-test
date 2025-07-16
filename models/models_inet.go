package models

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type RegisterRequest struct {
	Email          string `json:"email" validate:"required,email"`
	Username       string `json:"username" validate:"required,min=4,max=20"`
	Password       string `json:"password" validate:"required,min=6"`
	InlinePassword string `json:"inline_password" validate:"required"`
	BusinessType   string `json:"business_type" validate:"required"`
	NameWebsite    string `json:"name_website" validate:"required,min=2,max=30"`
}
