package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email        string
	Password     string
	Name         string
	Address      string
	Phone_number string
	Is_Active    bool
}

type UsersModel struct {
	DB *gorm.DB
}

func (um UsersModel) GetAll() ([]Users, error) {
	var result []Users
	err := um.DB.Find(&result).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return result, nil
}

func (um UsersModel) Get(_email, _password string) ([]Users, error) {
	var result []Users
	err := um.DB.Where(&Users{Email: _email, Password: _password}).First(&result).Error
	if err != nil {
		log.Println("Error on Query Email and Password", err.Error())
		return nil, err
	}
	return result, nil
}

func (um UsersModel) Create(newUser Users) (Users, error) {
	err := um.DB.Save(&newUser).Error
	if err != nil {
		fmt.Println("Error on Create", err.Error())
		return Users{}, err
	}
	return newUser, nil
}
