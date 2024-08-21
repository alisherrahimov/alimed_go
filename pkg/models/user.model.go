package models

// User struct

type User struct {
	Id          string `json:"id"`
	FName       string `json:"f_name" sql:"f_name"`
	LName       string `json:"l_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Gender      int    `json:"gender"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
