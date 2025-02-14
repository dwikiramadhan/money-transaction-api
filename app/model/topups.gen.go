// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTopup = "topups"

// Topup mapped from table <topups>
type Topup struct {
	ID            int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TopUpID       string     `gorm:"column:top_up_id;not null" json:"top_up_id"`
	UserID        string     `gorm:"column:user_id;not null" json:"user_id"`
	AmountTopUp   int32      `gorm:"column:amount_top_up;not null" json:"amount_top_up"`
	BalanceBefore int32      `gorm:"column:balance_before;not null" json:"balance_before"`
	BalanceAfter  int32      `gorm:"column:balance_after;not null" json:"balance_after"`
	CreatedDate   time.Time  `gorm:"column:created_date;not null" json:"created_date"`
	UpdatedDate   *time.Time `gorm:"column:updated_date" json:"updated_date"`
}

// TableName Topup's table name
func (*Topup) TableName() string {
	return TableNameTopup
}
