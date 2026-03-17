// Package auth
package auth

type SendOtpType string

const (
	SendOtpTypePhone SendOtpType = "phone"
	SendOtpTypeEmail SendOtpType = "email"
)

type SendOtpRequest struct {
	Identifier string      `json:"identifier" validate:"required"`
	Type       SendOtpType `json:"type" validate:"required,oneof=phone email"`
}
