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
