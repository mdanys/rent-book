package controllers

import (
	"fmt"
	"log"
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

func (uc UsersControl) Get(_username, _password string) ([]model.Users, error) {
	result, err := uc.Model.Get(_username, _password)
	if err != nil {
		fmt.Println("Error on Get", err.Error())
		return nil, err
	}
	return result, nil
}

func (uc UsersControl) GetById(_userId uint) (model.Users, error) {
	result, err := uc.Model.GetById(_userId)
	if err != nil {
		log.Println("Error on GetById", err.Error())
		return model.Users{}, err
	}
	return result, nil
}

func (uc UsersControl) Create(newUser model.Users) (model.Users, error) {
	result, err := uc.Model.Create(newUser)
	if err != nil {
		log.Println("Error on Create", err.Error())
		return model.Users{}, err
	}
	return result, nil
}

func (uc UsersControl) Edit(updatedUsers model.Users) (model.Users, error) {
	result, err := uc.Model.Edit(updatedUsers)
	if err != nil {
		log.Println("Error on Edit", err.Error())
		return model.Users{}, err
	}
	return result, nil
}

func (uc UsersControl) GetActive(_isActive bool) ([]model.Users, error) {
	result, err := uc.Model.GetActive(_isActive)
	if err != nil {
		fmt.Println("Error on GetActive", err.Error())
		return nil, err
	}
	return result, nil
}
