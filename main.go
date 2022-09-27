package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"rent-book/controllers"
	"rent-book/models"

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

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Users{})
	log.Println("Init migrate")
}

func callClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var isRunning bool = true
	var inputMenu int
	gconn, err := connectGorm()
	migrate(gconn)
	userModel := models.UsersModel{gconn}
	userControl := controllers.UsersControl{userModel}
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
			callClear()
			fmt.Println("\t--Login/Register--")
			fmt.Println("1. Login")
			fmt.Println("2. Register")
			fmt.Println("9. Back")
			fmt.Println("0. Main Menu")
			fmt.Print("Enter Your Input: ")
			fmt.Scanln(&inputMenu)
			switch inputMenu {
			case 1:
				var Get models.Users
				fmt.Println("Email:")
				fmt.Scanln(&Get)
				fmt.Println("Password:")
				fmt.Scanln(&Get)
			case 2:
				var newUsers models.Users
				fmt.Println("Input Name:")
				fmt.Scanln(&newUsers.Name)
				fmt.Println("Input Address:")
				fmt.Scanln(&newUsers.Address)
				fmt.Println("Input Phone:")
				fmt.Scanln(&newUsers.Phone_number)
				fmt.Println("Input Email:")
				fmt.Scanln(&newUsers.Email)
				fmt.Println("Input Password:")
				fmt.Scanln(&newUsers.Password)

				newUsers.Is_Active = true

				result, err := userControl.Create(newUsers)
				if err != nil {
					fmt.Println("Error on Adding User", err.Error())
				}
				// fmt.Println("Input :", newUsers)
				fmt.Println("\nRegistration User Success.", result)
			case 9:
				callClear()
				isRunning = true
			}
		case 9:
			callClear()
			isRunning = false
			fmt.Println("Thank you!")
		}
	}
}

func MenuLoginRegister() {

}
