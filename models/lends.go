package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Lends struct {
	gorm.Model
	Return_date time.Time
	Is_returned bool
	IDUser      uint
	IDBook      uint
}

type LendsModel struct {
	DB *gorm.DB
}

func (lm LendsModel) GetUserBorrow(_IDUser uint) ([]Books, error) {
	var result []Books
	err := lm.DB.Where(&Books{IDUser: _IDUser}).Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

// type BooksUsers struct {
// 	BooksID   int `gorm:"primaryKey"`
// 	UsersID   int `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	DeletedAt gorm.DeletedAt
// }
