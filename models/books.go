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

func (bm BooksModel) GetWhere(_title string) ([]Books, error) {
	var result []Books
	err := bm.DB.Where(&Books{Title: _title}).Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

func (bm BooksModel) GetUserBooks(_IDUser uint) ([]Books, error) {
	var result []Books
	err := bm.DB.Where(&Books{IDUser: _IDUser}).Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

func (bm BooksModel) GetBorrowed() ([]Books, error) {
	var result []Books
	err := bm.DB.Where(&Books{}).Find(&result).Error
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

func (bm BooksModel) Delete(deletedBooks Books) (Books, error) {
	err := bm.DB.Delete(&deletedBooks).Error
	if err != nil {
		fmt.Println("Error on Delete", err.Error())
		return Books{}, err
	}
	return deletedBooks, nil
}
