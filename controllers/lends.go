package controllers

import (
	"fmt"
	"rent-book/models"
)

type LendsControl struct {
	Model models.LendsModel
}

func (lc LendsControl) GetUserBookBorrow(_IDUser uint) ([]models.Lends, error) {
	result, err := lc.Model.GetUserBookBorrow(_IDUser)
	if err != nil {
		fmt.Println("Error GetUserBookBorrow")
		return nil, err
	}
	return result, nil
}

func (lc LendsControl) AddLend(_newLend models.Lends) (models.Lends, error) {
	result, err := lc.Model.AddLend(_newLend)
	if err != nil {
		fmt.Println("Error GetUserBookBorrow")
		return models.Lends{}, err
	}
	return result, nil
}
