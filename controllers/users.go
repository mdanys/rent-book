package controllers

import (
	model "rent-book/models"
)

type UsersControl struct {
	Model model.UsersModel
	Util  MyUtil
}

func (uc UsersControl) GetAll() ([]model.Users, error) {
	result, err := uc.Model.GetAll()
	if err != nil {
		uc.Util.ErrorMsg("Error on GetAll", "", err.Error())
		return nil, err
	}
	return result, nil
}

func (uc UsersControl) Get(_username, _password string) ([]model.Users, error) {
	result, err := uc.Model.Get(_username, _password)
	if err != nil {
		uc.Util.ErrorMsg("Error on Get", "", err.Error())
		return nil, err
	}
	return result, nil
}

func (uc UsersControl) GetById(_userId uint) (model.Users, error) {
	result, err := uc.Model.GetById(_userId)
	if err != nil {
		uc.Util.ErrorMsg("Error on GetById", "", err.Error())
		return model.Users{}, err
	}
	return result, nil
}

func (uc UsersControl) Create(newUser model.Users) (model.Users, error) {
	result, err := uc.Model.Create(newUser)
	if err != nil {
		uc.Util.ErrorMsg("Error on Create", "", err.Error())
		return model.Users{}, err
	}
	return result, nil
}

func (uc UsersControl) Edit(updatedUsers model.Users) (model.Users, error) {
	result, err := uc.Model.Edit(updatedUsers)
	if err != nil {
		uc.Util.ErrorMsg("Error on Edit", "", err.Error())
		return model.Users{}, err
	}
	return result, nil
}

func (uc UsersControl) GetActive(_isActive bool) ([]model.Users, error) {
	result, err := uc.Model.GetActive(_isActive)
	if err != nil {
		uc.Util.ErrorMsg("Error on GetActive", "", err.Error())
		return nil, err
	}
	return result, nil
}
