package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Lends struct {
	gorm.Model
	Return_date        time.Time
	Is_returned        bool
	Book_owner_id      uint
	Book_owner_email   string
	Book_owner_name    string
	Book_owner_address string
	Book_owner_phone   string
	Book_title         string
	Book_author        string
	IDUser             uint
	IDBook             uint
}

type LendsModel struct {
	DB *gorm.DB
}

//myUtil := controllers.myUtil{}

// Fungsi menampilkan list buku yang user pinjam
func (lm LendsModel) GetUserBookBorrow(_IDUser uint) ([]Lends, error) {
	var result []Lends
	err := lm.DB.Where("is_returned = ? AND id_user = ?", 0, _IDUser).Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

// Fungsi menambah peminjaman buku
func (lm LendsModel) AddLend(_newLend Lends) (Lends, error) {
	err := lm.DB.Save(&_newLend).Error
	if err != nil {
		return Lends{}, err
	}
	return _newLend, nil
}

// Fungsi Mengembalikan Buku
func (lm LendsModel) ReturnBook(_lendData Lends) (Lends, error) {
	err := lm.DB.Save(_lendData).Error
	if err != nil {
		return Lends{}, err
	}
	return _lendData, nil
}

// type BooksUsers struct {
// 	BooksID   int `gorm:"primaryKey"`
// 	UsersID   int `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	DeletedAt gorm.DeletedAt
// }
