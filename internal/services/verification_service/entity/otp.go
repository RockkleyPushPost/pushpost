package entity

import "time"

type OTPMessage struct {
	Email  string    `json:"email"`
	OTP    string    `json:"otp"`
	Expiry time.Time `json:"expiry"`
}
