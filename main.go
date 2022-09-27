package main

import (
	"fmt"
	"os"
	"os/exec"
	"rent-book/controllers"
	"rent-book/models"
	model "rent-book/models"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var isRunning bool = true
var isLoggedIn bool = false
var currentUser models.Users
var yesNo string

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
	var inputMenu int
	var next string
	gconn, err := connectGorm()
	migrate(gconn)
	userModel := models.UsersModel{DB: gconn}
	userControl := controllers.UsersControl{Model: userModel}
	if err != nil {
		fmt.Println("Can't connect to DB", err.Error())
	}

	for isRunning {
		callClear()

		if !isLoggedIn {
			fmt.Println("\t--Main Menu--")
			fmt.Println("1. Search Books")
			fmt.Println("2. List of Books")
			fmt.Println("3. Login/Register")
			fmt.Println("9. Exit")
			fmt.Print("Enter Your Input: ")
			fmt.Scanln(&inputMenu)
			switch inputMenu {
			case 1:
<<<<<<< Updated upstream
			case 2:
			case 3:
=======
				var Get model.Users
				fmt.Println("Email:")
				fmt.Scanln(&Get)
				fmt.Println("Password:")
				fmt.Scanln(&Get)
			case 2:
				var newUsers model.Users
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
>>>>>>> Stashed changes
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
					callClear()
					fmt.Println("\t--Login--")
					fmt.Println("Email:")
					fmt.Scanln(&currentUser.Email)
					fmt.Println("Password:")
					fmt.Scanln(&currentUser.Password)

					result, err := userControl.Get(currentUser.Email, currentUser.Password)
					if err != nil {
						Message("Login Gagal", "Email atau Password Salah.", err.Error())
					}
					if result != nil {
						isLoggedIn = true
						currentUser = result[0]
						fmt.Println("Login Sukses, Selamat Datang ", currentUser.Name)
						fmt.Println("Tekan Enter untuk ke Menu Member")
						fmt.Scanln(&next)
					}

				case 2:
					var newUsers models.Users
					callClear()
					fmt.Println("\t--Register--")
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
					fmt.Println("Terima Kasih")
					fmt.Println("Program keluar..")
					isRunning = false
				}
			case 9:
				fmt.Println("Terima Kasih")
				fmt.Println("Program keluar..")
				isRunning = false
			}
		} else {
			fmt.Println("\t--HOMEPAGE MEMBER--")
			fmt.Println("\tWelcome,", currentUser.Name)
			fmt.Println("1. Search Books")
			fmt.Println("2. List of Books")
			fmt.Println("3. Update Profile")
			fmt.Println("4. My Library")
			fmt.Println("===================")
			fmt.Println("9. Logout")
			fmt.Print("Enter Your Input: ")
			fmt.Scanln(&inputMenu)

			switch inputMenu {
			case 1:
			case 2:
			case 3:
			case 4:
			case 9:
				var isNotYesNo bool = true

				for isNotYesNo {
					fmt.Println("Apaah Anda yakin untuk Logout? (Y/N)")
					fmt.Scanln(&yesNo)
					if yesNo == "Y" || yesNo == "y" {
						isNotYesNo = false
						isLoggedIn = false
						callClear()
					} else if yesNo == "N" || yesNo == "n" {
						isNotYesNo = false
						callClear()
					} else {
						for isNotYesNo {

						}
					}
				}
			}
		case 9:
			callClear()
			isRunning = false
			fmt.Println("Thank you!")
		}
	}
}
<<<<<<< Updated upstream

func MenuLoginRegister() {

}

func Message(_title, _content, _detail string) {
	var next string
	fmt.Println("=== ", _title, " ===")
	fmt.Println(_content)
	fmt.Println(_detail)
	fmt.Println("Tekan Enter untuk melanjutkan.")
	fmt.Scanln(&next)
}
=======
>>>>>>> Stashed changes
