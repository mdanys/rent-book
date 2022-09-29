package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title       string
	Author      string
	Is_Borrowed bool
	Is_Deleted  bool
	IDUser      uint
	Lendss      []Lends `gorm:"foreignKey:IDBook;"`
	//Userss      []Users `gorm:"many2many:books_users;"`
	//Id_Category uint
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
func (bm BooksModel) GetAllAvailable() ([]Books, error) {
	var result []Books
	// err := bm.DB.Where(&Books{Is_Borrowed: false}).Find(&result).Error
	err := bm.DB.Not("is_borrowed = ?", 1).Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

func (bm BooksModel) GetWhere(_search, _column string) ([]Books, error) {
	var result []Books
	err := bm.DB.Where(_column+" LIKE ?", "%"+_search+"%").Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

func (bm BooksModel) GetUserBooks(_IDUser uint) ([]Books, error) {
	var result []Books
	err := bm.DB.Where("id_user = ?", _IDUser).Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

func (bm BooksModel) GetById(_IDBook uint) (Books, error) {
	var result Books
	err := bm.DB.First(&result, _IDBook).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return Books{}, err
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

func (bm BooksModel) Delete(deletedBooks Books) (Books, error) {
	err := bm.DB.Delete(&deletedBooks).Error
	if err != nil {
		fmt.Println("Error on Delete", err.Error())
		return Books{}, err
	}
	return deletedBooks, nil
}
