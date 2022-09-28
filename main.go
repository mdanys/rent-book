package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"rent-book/controllers"
	"rent-book/models"
	model "rent-book/models"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var showError bool = true
var isRunning bool = true
var isLoggedIn bool = false
var currentUser models.Users
var currentBook models.Books
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
	db.AutoMigrate(&model.Books{})
}

func main() {
	var inputMenu int
	gconn, err := connectGorm()
	migrate(gconn)
	userModel := models.UsersModel{DB: gconn}
	userControl := controllers.UsersControl{Model: userModel}
	bookModel := models.BooksModel{DB: gconn}
	bookControl := controllers.BooksControl{Model: bookModel}
	// lendModel := models.LendsModel{DB: gconn}
	// lendControl := controllers.LendsControl{Model: lendModel}

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
				callClear()
				fmt.Println("\t--Find Books--")
				fmt.Println("1. Find Books by Title")
				fmt.Println("2. Find Books by Author")
				fmt.Println("9. Back")
				fmt.Println("0. Main Menu")
				fmt.Print("Enter Your Input: ")
				fmt.Scanln(&inputMenu)
				switch inputMenu {
				case 1:
					callClear()
					fmt.Println("\t--Find Books by Title--")
					fmt.Println("Please enter book's title :")
					fmt.Scanln(&currentBook.Title)

					result, err := bookControl.GetWhere(currentBook.Title)
					if err != nil {
						ErrorMsg(showError, "Searching Book Failed", "Book not found or failed to retrieve data.", err.Error())
					}

					fmt.Println("Tabel Pencarian Buku")
					fmt.Println("==================================")
					fmt.Printf("%4s | %5s | %15s | %15s | %15s |\n", "No", "Book Id", "Title", "Author", "Status")

					if result != nil {
						//i := 1
						// for _, value := range result {
						// 	//fmt.Printf("%4s | %5s | %15s | %15s | %15s |\n", i, value.ID, value.Title, value.Author, value."Status")
						// 	i++
						// }
					} else {
						fmt.Println("\n\t\\tt Book Title not Found")
					}

					fmt.Println("9. Back")
					fmt.Println("0. Main Menu")
					fmt.Print("Enter Your Input: ")
					fmt.Scanln(&inputMenu)
				case 2:
				}
			case 2:
			case 3:
				isLoginRegisterMenu := true

				for isLoginRegisterMenu {

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
							ErrorMsg(showError, "Login Failed", "Wrong Email or Password. Please try again.", err.Error())
						}
						if result != nil {
							currentUser = result[0]
							if currentUser.Is_Active == true {
								isLoggedIn = true
								msg := "Login Sukses, Selamat Datang " + currentUser.Name
								Message("Login Success", msg, "Move to Homepage Member")
								isLoginRegisterMenu = false
							} else {
								msg := "Login Failed, Account is Deactivated"
								Message("Login Failed", msg, "Please register again")
							}
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
							ErrorMsg(showError, "User Registration Failed", "Please try again.", err.Error())
						}

						msg := "User " + result.Email + " success registered."
						Message("User Registration Success", msg, "")
					case 9:
						Message("Kembali", "Ke Menu Sebelumnya...", "")
						isLoginRegisterMenu = false
					}
				}
			case 9:
				Message("Terima Kasih", "Program Keluar ...", "")
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
				updateProfile := true
				for updateProfile {
					callClear()

					fmt.Println("\t--Update Profile--")
					fmt.Println("=========================")
					fmt.Printf("Name: %s\nAddress: %s\nNumber Phone: %s\nEmail: %s", currentUser.Name, currentUser.Address, currentUser.Phone_number, currentUser.Email)
					fmt.Println("\n=========================")

					fmt.Println("1. Edit Name")
					fmt.Println("2. Edit Address")
					fmt.Println("3. Edit Phone Number")
					fmt.Println("4. Change Password")
					fmt.Println("5. Deactivate Account")
					fmt.Println("========================")
					fmt.Println("9. Back")
					fmt.Println("0. Main Menu")
					fmt.Print("Enter Your Input: ")
					fmt.Scanln(&inputMenu)
					switch inputMenu {
					case 1:
						callClear()
						fmt.Println("\t--Edit Name--")
						fmt.Println("================")
						fmt.Printf("Name: %s", currentUser.Name)
						fmt.Println("\n================")
						fmt.Println("Input New Name :")
						fmt.Scanln(&currentUser.Name)
						resUpdateUser, err := userControl.Edit(currentUser)
						if err != nil {
							ErrorMsg(showError, "Error on Updating Name", "Please try again.", err.Error())
						}
						Message("Success", "Updating New Name Success", resUpdateUser)

						updateProfile = false
					case 2:
						callClear()
						fmt.Println("\t--Edit Address--")
						fmt.Println("================")
						fmt.Printf("Address: %s", currentUser.Address)
						fmt.Println("\n================")
						fmt.Println("Input New Address :")
						fmt.Scanln(&currentUser.Address)
						resUpdateUser, err := userControl.Edit(currentUser)
						if err != nil {
							ErrorMsg(showError, "Error on Updating Address", "Please try again.", err.Error())
						}
						Message("Success", "Updating New Address Success", resUpdateUser)

						updateProfile = false
					case 3:
						callClear()
						fmt.Println("\t--Edit Phone Number--")
						fmt.Println("================")
						fmt.Printf("Address: %s", currentUser.Phone_number)
						fmt.Println("\n================")
						fmt.Println("Input New Phone Number :")
						fmt.Scanln(&currentUser.Phone_number)
						resUpdateUser, err := userControl.Edit(currentUser)
						if err != nil {
							ErrorMsg(showError, "Error on Updating Phone Number", "Please try again.", err.Error())
						}
						Message("Success", "Updating New Phone Number Success", resUpdateUser)

						updateProfile = false
					case 4:
						callClear()
						var oldPassword, newPassword, rePassword string
						isChangePassword := true
						for isChangePassword {
							fmt.Println("\t--Change Password--")
							fmt.Println("================")
							fmt.Println("Input Your Old Password")
							fmt.Scanln(&oldPassword)
							fmt.Println("Input Your New Password")
							fmt.Scanln(&newPassword)
							fmt.Println("Re-enter Your New Password")
							fmt.Scanln(&rePassword)
							if currentUser.Password == oldPassword {
								if newPassword == rePassword {
									currentUser.Password = newPassword
									resUpdateUser, err := userControl.Edit(currentUser)
									if err != nil {
										ErrorMsg(showError, "Error on Updating Password", "Please try again.", err.Error())
									}
									Message("Success", "Updating New Password", resUpdateUser)

									isChangePassword = false
									updateProfile = false
									isLoggedIn = false
									currentUser = model.Users{}
								} else {
									Message("Failed", "New Password is not match", "Please try again")
									callClear()
								}
							} else {
								Message("Failed", "Old Password is not correct", "Please try again")
								isChangePassword = false
							}
						}
					case 5:
						callClear()
						fmt.Println("\t--Status Account--")
						fmt.Println("================")
						fmt.Printf("Status: %t", currentUser.Is_Active)
						fmt.Println("\n================")
						isDeactivateAcc := true
						for isDeactivateAcc {
							fmt.Println("\nAre you sure want to Deactivate Your Account? (Y/N)")
							fmt.Scanln(&yesNo)
							if yesNo == "Y" || yesNo == "y" {
								currentUser.Is_Active = false
								resUpdateUser, err := userControl.Edit(currentUser)
								if err != nil {
									ErrorMsg(showError, "Error on Deactivating Account", "Please try again.", err.Error())
								}
								Message("Success", "Deactivate Account Success", resUpdateUser)
								isDeactivateAcc = false
								updateProfile = false
								isLoggedIn = false
								currentUser = model.Users{}
								callClear()
							} else if yesNo == "N" || yesNo == "n" {
								isDeactivateAcc = false
								updateProfile = false
								callClear()
							} else {
								isDeactivateAcc = true
							}
						}
					case 9:
					case 0:
					}
				}
			case 4:
				callClear()
				fmt.Println("\t--My Library--")
				fmt.Println("1. My Books")
				fmt.Println("2. Borrow")

				fmt.Println("===================")
				fmt.Println("9. Back")
				fmt.Println("0. Main Menu")
				fmt.Print("Enter Your Input: ")
				fmt.Scanln(&inputMenu)

				switch inputMenu {
				case 1:
					callClear()
					fmt.Println("\t--My Books--")
					fmt.Println("1. Add Book")
					fmt.Println("2. List My Book")

					fmt.Println("===================")
					fmt.Println("9. Back")
					fmt.Println("0. Main Menu")
					fmt.Print("Enter Your Input: ")
					fmt.Scanln(&inputMenu)
					switch inputMenu {
					case 1:
						var newBook models.Books
						callClear()
						fmt.Println("\t--Add Book--")
						fmt.Println("Input Title:")
						fmt.Scanln(&newBook.Title)
						fmt.Println("Input Author:")
						fmt.Scanln(&newBook.Author)

						newBook.Is_Borrowed = false
						newBook.Is_Deleted = false
						newBook.IDUser = currentUser.ID

						result, err := bookControl.Add(newBook)
						if err != nil {
							ErrorMsg(showError, "Add Book Failed", "Failed to Add Book. Please try again.", err.Error())
						}

						Message("Add Book Success", "Book "+result.Title+" Added.", "")
					case 2:
						listMyBookMenu := true
						for listMyBookMenu {

							callClear()
							var inputBookNumber int

							fmt.Println("\t--List My Book--")

							result, err := bookControl.GetUserBooks(currentBook.ID)
							if err != nil {
								ErrorMsg(showError, "Failed Retrieve List Books", "Please try again.", err.Error())
							}

							fmt.Println("List My Book Table")
							fmt.Println("==================================")
							fmt.Printf("%4s | %5s | %15s | %15s | %15s |\n", "No", "Book Id", "Title", "Author", "Status")

							if result != nil {
								i := 1
								var status string
								for _, value := range result {
									if value.Is_Borrowed {
										status = "Not Available"
									} else {
										status = "Available"
									}
									fmt.Printf("%4d | %5d | %15s | %15s | %15s |\n", i, value.ID, value.Title, value.Author, status)
									i++
								}
							} else {
								fmt.Println("\n\t\\tt Book Title not Found")
							}

							fmt.Println("\n==============================")
							fmt.Println("1. Choose Number Edit Book Data")
							fmt.Println("9. Back")
							fmt.Println("0. Main Menu")
							fmt.Print("Enter Your Input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("== Choose Number Edit Book Data ==")
								fmt.Scanln(&inputBookNumber)

								var targetedBook model.Books = result[inputBookNumber-1]

								callClear()
								fmt.Println("List My Book Table")
								fmt.Println("==================================")
								fmt.Printf("%4s | %5s | %15s | %15s | %15s |\n", "No", "Book Id", "Title", "Author", "Status")

								var status string
								if targetedBook.Is_Borrowed {
									status = "Not Available"
								} else {
									status = "Available"
								}
								fmt.Printf("%4d | %5d | %15s | %15s | %15s |\n", 1, targetedBook.ID, targetedBook.Title, targetedBook.Author, status)

								fmt.Println("\n==============================")
								fmt.Println("1. Edit Title")
								fmt.Println("2. Edit Author")
								fmt.Println("3. Delete Book")
								fmt.Println("9. Back")
								fmt.Println("0. Main Menu")
								fmt.Print("Enter Your Input: ")
								fmt.Scanln(&inputMenu)
								switch inputMenu {
								case 1:
									fmt.Println("\t--Edit Title--")
									fmt.Println("Input Title :")
									fmt.Scanln(&targetedBook.Title)
									resUpdateBook, err := bookControl.Edit(targetedBook)
									if err != nil {
										ErrorMsg(showError, "Error on Updating Book", "Please try again.", err.Error())
									}
									Message("Success", "Adding Book Success", resUpdateBook)

									// kembali ke menu list my book
									listMyBookMenu = false
								case 2:
									fmt.Println("\t--Edit Author--")
									fmt.Println("Input Author :")
									fmt.Scanln(&targetedBook.Author)
									resUpdateBook, err := bookControl.Edit(targetedBook)
									if err != nil {
										ErrorMsg(showError, "Error on Updating Book", "Please try again", err.Error())
									}
									Message("Success", "Adding Book Success", resUpdateBook)

									// kembali ke menu list my book
									listMyBookMenu = false
								case 3:
									fmt.Println("\t--Delete Book--")

									isNotYesNo := true
									for isNotYesNo {
										fmt.Println("Are you sure to delete this? (y/n)")
										fmt.Scanln(&yesNo)
										if yesNo == "Y" || yesNo == "y" {
											resultDelete, err := bookControl.Delete(targetedBook)
											if err != nil {
												ErrorMsg(showError, "Failed", "Deleting Book Failed", resultDelete)
												fmt.Println("", err.Error())
											}
											Message("Success", "Deleting Book Success", resultDelete)

											isNotYesNo = false
											listMyBookMenu = false
											callClear()
										} else if yesNo == "N" || yesNo == "n" {
											isNotYesNo = false
											listMyBookMenu = false
											callClear()
										} else {
											isNotYesNo = true
										}
									}
								case 9:
								case 0:
								}

							case 9:
							case 0:

							}
						}
					}
				case 2:
				}
			case 9:
				var isNotYesNo bool = true

				for isNotYesNo {
					fmt.Println("Are you sure to Logout? (Y/N)")
					fmt.Scanln(&yesNo)
					if yesNo == "Y" || yesNo == "y" {
						isNotYesNo = false
						isLoggedIn = false
						callClear()
					} else if yesNo == "N" || yesNo == "n" {
						isNotYesNo = false
						callClear()
					} else {
						isNotYesNo = true
					}
				}
			}
		}
	}
}

func MenuLoginRegister() {

}

func Message(_title, _content, _detail interface{}) {
	var next string
	fmt.Println("=== ", _title, " ===")
	fmt.Println(_content)
	fmt.Println(_detail)
	fmt.Println("Tekan Enter untuk melanjutkan.")
	fmt.Scanln(&next)
}

func ErrorMsg(_isShow bool, _title, _content, _detail interface{}) {
	var next string
	if _isShow {
		fmt.Println("=== SHOW ERROR LOG ===")
		fmt.Println("=== ", _title, " ===")
		fmt.Println(_content)
		log.Println(_detail)
	} else {
		fmt.Println("=== ", _title, " ===")
		fmt.Println(_content)
	}
	fmt.Println("\nTekan Enter untuk melanjutkan.")
	fmt.Scanln(&next)
}
