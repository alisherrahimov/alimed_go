package models

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type LoginByEmailRequest struct {
	Email string `json:"email"`
}

type VerifyOTPRequest struct {
	OTP         int    `json:"otp"`
	PhoneNumber string `json:"phone_number"`
}

type AskSomeDataRequest struct {
	PhoneNumber string `json:"phone_number"`
	FName       string `json:"f_name"`
	LName       string `json:"l_name"`
	Gender      int    `json:"gender"`
	Age         int    `json:"age"`
}
