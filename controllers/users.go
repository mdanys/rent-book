package controllers

import (
	"fmt"
	model "rent-book/models"
)

type UsersControl struct {
	Model model.UsersModel
}

func (uc UsersControl) GetAll() ([]model.Users, error) {
	result, err := uc.Model.GetAll()
	if err != nil {
		fmt.Println("Error on GetAll", err.Error())
		return nil, err
	}
	return result, nil
}

func (uc UsersControl) Get() ([]model.Users, error) {
	result, err := uc.Model.Get()
	if err != nil {
		fmt.Println("Error on Get", err.Error())
		return nil, err
	}
	return result, nil
}

func (uc UsersControl) Create() (model.Users, error) {
	result, err := uc.Model.Create()
	if err != nil {
		fmt.Println("Error on Create", err.Error())
		return model.Users{}, err
	}
	return result, nil
}
