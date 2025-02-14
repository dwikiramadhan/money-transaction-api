package entity

import "time"

type SignUpReq struct {
	FirstName   *string `gorm:"column:first_name" json:"first_name"`
	LastName    *string `gorm:"column:last_name" json:"last_name"`
	PhoneNumber *string `gorm:"column:phone_number" json:"phone_number"`
	Address     *string `gorm:"column:address" json:"address"`
	Pin         *string `gorm:"column:pin" json:"pin"`
}

type SignUpReResp struct {
	UserID      string    `gorm:"column:user_id;not null" json:"user_id"`
	FirstName   *string   `gorm:"column:first_name" json:"first_name"`
	LastName    *string   `gorm:"column:last_name" json:"last_name"`
	PhoneNumber *string   `gorm:"column:phone_number" json:"phone_number"`
	Address     *string   `gorm:"column:address" json:"address"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
}
