package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/rent-book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	var isRunning bool = true
	var inputMenu int

	for isRunning {
		fmt.Println("\t--Main Menu--")
		fmt.Println("1. Search Books")
		fmt.Println("2. List of Books")
		fmt.Println("3. Login/Register")
		fmt.Println("9. Exit")
		fmt.Print("Enter Your Input: ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
		case 2:
		case 3:
		}
	}
}
