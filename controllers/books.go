package controllers

import (
	"fmt"
	"log"
	model "rent-book/models"
)

type BooksControl struct {
	Model model.BooksModel
}

func (bc BooksControl) GetAll() ([]model.Books, error) {
	result, err := bc.Model.GetAll()
	if err != nil {
		fmt.Println("Error on GetAll", err.Error())
		return nil, err
	}
	return result, nil
}

func (bc BooksControl) GetAllAvailable() ([]model.Books, error) {
	result, err := bc.Model.GetAllAvailable()
	if err != nil {
		fmt.Println("Error on GetAll", err.Error())
		return nil, err
	}
	return result, nil
}

func (bc BooksControl) GetWhere(_title string) ([]model.Books, error) {
	result, err := bc.Model.GetWhere(_title)
	if err != nil {
		fmt.Println("Error on GetWhere", err.Error())
		return nil, err
	}
	return result, nil
}

func (bc BooksControl) GetUserBooks(_IDUser uint) ([]model.Books, error) {
	result, err := bc.Model.GetUserBooks(_IDUser)
	if err != nil {
		fmt.Println("Error on GetWhere", err.Error())
		return nil, err
	}
	return result, nil
}

func (bc BooksControl) GetById(_IDBook uint) (model.Books, error) {
	result, err := bc.Model.GetById(_IDBook)
	if err != nil {
		fmt.Println("Error on GetWhere", err.Error())
		return model.Books{}, err
	}
	return result, nil
}

func (bc BooksControl) Add(newBooks model.Books) (model.Books, error) {
	result, err := bc.Model.Add(newBooks)
	if err != nil {
		log.Println("Error on Add", err.Error())
		return model.Books{}, err
	}
	return result, nil
}

func (bc BooksControl) Edit(updatedBooks model.Books) (model.Books, error) {
	result, err := bc.Model.Edit(updatedBooks)
	if err != nil {
		log.Println("Error on Edit", err.Error())
		return model.Books{}, err
	}
	return result, nil
}

func (bc BooksControl) Delete(deletedBooks model.Books) (model.Books, error) {
	result, err := bc.Model.Delete(deletedBooks)
	if err != nil {
		log.Println("Error on Delete", err.Error())
		return model.Books{}, err
	}
	return result, nil
}
