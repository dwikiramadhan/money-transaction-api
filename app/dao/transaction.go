package dao

import (
	"fmt"
	"time"

	"github.com/dwikiramadhan/money-transaction-api/app/entity"
	"github.com/dwikiramadhan/money-transaction-api/app/model"
	"github.com/google/uuid"
)

func Topup(record *entity.TopupReq) (result *model.Topup, RowAffected int64, err error) {
	latestBalance, _, err := GetLatestBalance()
	if err != nil {
		return nil, -1, fmt.Errorf("Error on get latest balance, %w", err)
	}

	totalBalanceAfter := latestBalance.BalanceAfter + record.Amount

	balance := &model.Topup{
		TopUpID:       uuid.NewString(),
		UserID:        "ff638c9c-e23e-4dc5-be1d-783d62320f39",
		AmountTopUp:   record.Amount,
		BalanceBefore: latestBalance.BalanceAfter,
		BalanceAfter:  totalBalanceAfter,
		CreatedDate:   time.Now(),
	}
	db := DB.Debug().Create(balance)

	if err = db.Error; err != nil {
		// Return only error.
		return nil, -1, ErrInsertFailed
	}

	// This query returns nothing.
	return balance, db.RowsAffected, nil

}

func Payment(record *entity.PayReq) (result *model.Payment, RowAffected int64, err error) {
	latestBalance, _, err := GetLatestBalance()
	if err != nil {
		return nil, -1, fmt.Errorf("Error on get latest balance, %w", err)
	}

	totalBalanceAfter := latestBalance.BalanceAfter + record.Amount

	payment := &model.Payment{
		PaymentID:     uuid.NewString(),
		UserID:        "ff638c9c-e23e-4dc5-be1d-783d62320f39",
		Amount:        &record.Amount,
		Remarks:       &record.Remarks,
		BalanceBefore: latestBalance.BalanceAfter,
		BalanceAfter:  totalBalanceAfter,
		CreatedDate:   time.Now(),
	}
	db := DB.Debug().Create(payment)

	if err = db.Error; err != nil {
		// Return only error.
		return nil, -1, ErrInsertFailed
	}

	// This query returns nothing.
	return payment, db.RowsAffected, nil

}

func Transfer(record *entity.TransferReq) (result *model.Transfer, RowAffected int64, err error) {
	latestBalance, _, err := GetLatestBalance()
	if err != nil {
		return nil, -1, fmt.Errorf("Error on get latest balance, %w", err)
	}

	totalBalanceAfter := latestBalance.BalanceAfter + record.Amount

	payment := &model.Transfer{
		TransferID:    uuid.NewString(),
		UserID:        "ff638c9c-e23e-4dc5-be1d-783d62320f39",
		Amount:        record.Amount,
		Remarks:       &record.Remarks,
		BalanceBefore: latestBalance.BalanceAfter,
		BalanceAfter:  totalBalanceAfter,
		CreatedDate:   time.Now(),
	}
	db := DB.Debug().Create(payment)

	if err = db.Error; err != nil {
		// Return only error.
		return nil, -1, ErrInsertFailed
	}

	// This query returns nothing.
	return payment, db.RowsAffected, nil

}

func GetLatestBalance() (result *model.Topup, rows int64, err error) {

	resultOrm := DB.Model(&model.Topup{})
	resultOrm.Debug().Last(&result).Count(&rows)
	if err = resultOrm.Error; err != nil {
		// Return empty object and error.
		err = ErrNotFound
		return nil, -1, err
	}
	// Return query result.
	return result, rows, nil
}
