package dao

import (
	"time"

	"github.com/dwikiramadhan/money-transaction-api/app/entity"
	"github.com/dwikiramadhan/money-transaction-api/app/model"
	"github.com/google/uuid"
)

func CreateUser(record *entity.SignUpReq) (result *model.User, RowAffected int64, err error) {

	user := &model.User{
		UserID:      uuid.NewString(),
		FirstName:   record.FirstName,
		LastName:    record.LastName,
		PhoneNumber: record.PhoneNumber,
		Address:     record.Address,
		Pin:         record.Pin,
		CreatedAt:   time.Now(),
	}
	db := DB.Debug().Create(user)

	if err = db.Error; err != nil {
		// Return only error.
		return nil, -1, ErrInsertFailed
	}

	// This query returns nothing.
	return user, db.RowsAffected, nil

}

func GetUserByPhoneNumber(phone_number string) (result *model.User, rows int64, err error) {

	resultOrm := DB.Model(&model.User{})
	resultOrm.Debug().First(&result, "phone_number = ?", phone_number).Count(&rows)
	if err = resultOrm.Error; err != nil {
		// Return empty object and error.
		err = ErrNotFound
		return nil, -1, err
	}
	// Return query result.
	return result, rows, nil
}
