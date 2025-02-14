package entity

type LoginReq struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	PIN         string `json:"pin" validate:"required"`
}
