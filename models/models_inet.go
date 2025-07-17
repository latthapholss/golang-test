package models

import "gorm.io/gorm"

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
	LineID         string `json:"line_id"`
	Phone          string `json:"phone" validate:"required"`
	BusinessType   string `json:"business_type" validate:"required"`
	NameWebsite    string `json:"name_website" validate:"required,min=2,max=30"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data       []DogsRes `json:"data"`
	Name       string    `json:"name"`
	Count      int       `json:"count"`
	SumRed     int       `json:"sum_red"`
	SumGreen   int       `json:"sum_green"`
	SumPink    int       `json:"sum_pink"`
	SumNoColor int       `json:"sum_no_color"`
}

// company struct
type Company struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type CompanyRes struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type CompanyReq struct {
	Name    string `json:"name" validate:"required,min=3,max=32"`
	Address string `json:"address" validate:"required,min=3,max=64"`
	Phone   string `json:"phone" validate:"required,min=10,max=15"`
	Email   string `json:"email" validate:"required,email,min=3,max=32"`
}
