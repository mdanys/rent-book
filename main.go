package main

import (
	"fmt"
	"os"
	"os/exec"
	model "rent-book/models"

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

func callClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Users{})
}

func main() {
	var isRunning bool = true
	var inputMenu int
	gconn, err := connectGorm()
	migrate(gconn)
	// userModel := models.UsersModel{gconn}
	// userControl := controllers.UsersControl{userModel}
	if err != nil {
		fmt.Println("Can't connect to DB", err.Error())
	}

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

		case 9:
			callClear()
			isRunning = false
			fmt.Println("Thank you!")
		}
	}
}
