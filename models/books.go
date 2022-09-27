package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title       string
	Author      string
	is_Borrowed bool
	is_Deleted  bool
	Id_User     uint
	Id_Category uint
}

type BooksModel struct {
	DB *gorm.DB
}

func (bm BooksModel) GetAll() ([]Books, error) {
	var result []Books
	err := bm.DB.Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

func (bm BooksModel) Add(newBooks Books) (Books, error) {
	err := bm.DB.Save(&newBooks).Error
	if err != nil {
		fmt.Println("Error on Create", err.Error())
		return Books{}, err
	}
	return newBooks, nil
}

func (bm BooksModel) Edit(updatedBooks Books) (Books, error) {
	err := bm.DB.Save(&updatedBooks).Error
	if err != nil {
		fmt.Println("Error on Edit", err.Error())
		return Books{}, err
	}
	return updatedBooks, nil
}
