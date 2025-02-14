package entity

import "time"

type TopupReq struct {
	Amount int32 `json:"amount" validate:"required"`
}
type TopupResp struct {
	TopUpID       string    `gorm:"column:top_up_id;not null" json:"top_up_id"`
	AmountTopUp   int32     `gorm:"column:amount_top_up;not null" json:"amount_top_up"`
	BalanceBefore int32     `gorm:"column:balance_before;not null" json:"balance_before"`
	BalanceAfter  int32     `gorm:"column:balance_after;not null" json:"balance_after"`
	CreatedDate   time.Time `gorm:"column:created_date;not null" json:"created_date"`
}

type PayReq struct {
	Amount  int32  `json:"amount" validate:"required"`
	Remarks string `json:"remarks"`
}

type PaymentResp struct {
	PaymentID     string    `gorm:"column:payment_id;not null" json:"payment_id"`
	Amount        *int32    `gorm:"column:amount" json:"amount"`
	Remarks       *string   `gorm:"column:remarks" json:"remarks"`
	BalanceBefore int32     `gorm:"column:balance_before;not null" json:"balance_before"`
	BalanceAfter  int32     `gorm:"column:balance_after;not null" json:"balance_after"`
	CreatedDate   time.Time `gorm:"column:created_date;not null" json:"created_date"`
}

type TransferReq struct {
	TargetUser string `json:"target_user" validate:"required"`
	Amount     int32  `json:"amount" validate:"required"`
	Remarks    string `json:"remarks"`
}

type TransferResp struct {
	TransferID    string    `gorm:"column:transfer_id;not null" json:"transfer_id"`
	Amount        int32     `gorm:"column:amount;not null" json:"amount"`
	Remarks       *string   `gorm:"column:remarks" json:"remarks"`
	BalanceBefore int32     `gorm:"column:balance_before;not null" json:"balance_before"`
	BalanceAfter  int32     `gorm:"column:balance_after;not null" json:"balance_after"`
	CreatedDate   time.Time `gorm:"column:created_date;not null" json:"created_date"`
}
