package controllers

import (
	"rent-book/models"
)

type LendsControl struct {
	Model models.LendsModel
	Util  MyUtil
}

func (lc LendsControl) GetUserBookBorrow(_IDUser uint) ([]models.Lends, error) {
	result, err := lc.Model.GetUserBookBorrow(_IDUser)
	if err != nil {
		//fmt.Println("Error GetUserBookBorrow")
		lc.Util.ErrorMsg("Error GetUserBorrow", "", err.Error())
		return nil, err
	}
	return result, nil
}

func (lc LendsControl) AddLend(_newLend models.Lends) (models.Lends, error) {
	result, err := lc.Model.AddLend(_newLend)
	if err != nil {
		lc.Util.ErrorMsg("Error GetUserBookBorrow", "", err.Error())
		return models.Lends{}, err
	}
	return result, nil
}
