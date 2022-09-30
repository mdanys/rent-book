package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"rent-book/controllers"
	"rent-book/models"
	model "rent-book/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var showError bool = false
var isRunning bool = false
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
	db.AutoMigrate(&model.Lends{})
}

func main() {
	var inputMenu int
	gconn, errDB := connectGorm()
	migrate(gconn)
	userModel := models.UsersModel{DB: gconn}
	userControl := controllers.UsersControl{Model: userModel}
	bookModel := models.BooksModel{DB: gconn}
	bookControl := controllers.BooksControl{Model: bookModel}
	lendModel := models.LendsModel{DB: gconn}
	lendControl := controllers.LendsControl{Model: lendModel, Util: controllers.MyUtil{}}
	//utilCtl := controllers.MyUtil{}

	if errDB != nil {
		ErrorMsg(showError, "Failed Connect to Database", "Please check your database connection.", errDB.Error())
		isRunning = false
	} else {
		isRunning = true
	}

	// Auto Login
	// resultAutoLogin, _ := userControl.Get("doi@gmail.com", "123")
	// currentUser = resultAutoLogin[0]
	// isLoggedIn = true

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
				fmt.Println("\t--Search Books--")
				fmt.Println("1. Find Books by Title")
				fmt.Println("2. Find Books by Author")
				fmt.Println("9. Back")
				fmt.Print("Enter Your Input: ")
				fmt.Scanln(&inputMenu)
				switch inputMenu {
				case 1:
					callClear()
					fmt.Println("\t--Find Books by Title--")
					fmt.Println("Please enter book's title :")
					fmt.Scanln(&currentBook.Title)

					result, err := bookControl.GetWhere(currentBook.Title, "title")
					if err != nil {
						ErrorMsg(showError, "Searching Book Failed", "Book not found or failed to retrieve data.", err.Error())
					}
					var title string = fmt.Sprintf("%4s | %s | %s | %25s | %13s |\n", "No", "Book Id", fmt.Sprintf("%*s", 25, "Title"), "Author", "Status")
					rowLen := len(title)

					fmt.Println("== List of Book ==")
					fmt.Print(title)
					for i := 0; i < rowLen; i++ {
						fmt.Print("=")
					}
					fmt.Print("\n")
					if result != nil {
						i := 1
						for _, value := range result {
							fmt.Printf("%4d | %s | %25s | %25s | %s |\n", i, fmt.Sprintf("%*d", 7, value.ID), value.Title, value.Author, fmt.Sprintf("%*s", 13, BorrowedStatus(value.Is_Borrowed)))
							i++
						}
					} else {
						fmt.Println("\n\t\\tt Book Title not Found")
					}
					Message("Done", "", "")
				case 2:
					callClear()
					fmt.Println("\t--Find Books by Author--")
					fmt.Println("Please enter book's author :")
					fmt.Scanln(&currentBook.Author)

					result, err := bookControl.GetWhere(currentBook.Author, "author")
					if err != nil {
						ErrorMsg(showError, "Searching Book Failed", "Book not found or failed to retrieve data.", err.Error())
					}
					var author string = fmt.Sprintf("%4s | %s | %s | %25s | %13s |\n", "No", "Book Id", fmt.Sprintf("%*s", 25, "Title"), "Author", "Status")
					rowLen := len(author)

					fmt.Println("== List of Book ==")
					fmt.Print(author)
					for i := 0; i < rowLen; i++ {
						fmt.Print("=")
					}
					fmt.Print("\n")
					if result != nil {
						i := 1
						for _, value := range result {
							fmt.Printf("%4d | %s | %25s | %25s | %s |\n", i, fmt.Sprintf("%*d", 7, value.ID), value.Title, value.Author, fmt.Sprintf("%*s", 13, BorrowedStatus(value.Is_Borrowed)))
							i++
						}
					} else {
						fmt.Println("\n\t\\tt Book Title not Found")
					}
					Message("Done", "", "")
				case 9:
					Message("Kembali", "Ke Menu Sebelumnya...", "")
					isRunning = false
				}
			case 2:
				callClear()
				result, err := bookControl.GetAll()
				if err != nil {
					ErrorMsg(showError, "Searching Book Failed", "Book not found or failed to retrieve data.", err.Error())
				}
				var title string = fmt.Sprintf("%4s | %s | %s | %25s | %5s |\n", "No", "Book Id", fmt.Sprintf("%*s", 25, "Title"), "Author", "ID Owner")
				rowLen := len(title)

				fmt.Println("== List of Book ==")
				fmt.Print(title)
				for i := 0; i < rowLen; i++ {
					fmt.Print("=")
				}
				fmt.Print("\n")
				if result != nil {
					i := 1
					for _, value := range result {
						fmt.Printf("%4d | %s | %25s | %25s | %s |\n", i, fmt.Sprintf("%*d", 7, value.ID), value.Title, value.Author, fmt.Sprintf("%*d", 8, value.IDUser))
						i++
					}
				} else {
					fmt.Println("\n\t\\tt Book Title not Found")
				}
				Message("Done", "", "")
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
					case 0:
						Message("Kembali", "Ke Menu Sebelumnya...", "")
						isLoginRegisterMenu = false
					}
				}
			case 9:
				Message("Terima Kasih", "Program Keluar ...", "")
				isRunning = false
			}
		} else {
			homepageMemberMenu := true
			for homepageMemberMenu {
				callClear()
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
					callClear()
					fmt.Println("\t--Search Books--")
					fmt.Println("1. Find Books by Title")
					fmt.Println("2. Find Books by Author")
					fmt.Println("9. Back")
					fmt.Print("Enter Your Input: ")
					fmt.Scanln(&inputMenu)
					switch inputMenu {
					case 1:
						isFindTitle := true

						for isFindTitle {
							callClear()
							fmt.Println("\t--Find Books by Title--")
							fmt.Println("Please enter book's title :")
							fmt.Scanln(&currentBook.Title)

							result, err := bookControl.GetWhere(currentBook.Title, "title")
							if err != nil {
								ErrorMsg(showError, "Searching Book Failed", "Book not found or failed to retrieve data.", err.Error())
							}

							var title string = fmt.Sprintf("%4s | %s | %s | %25s | %13s |\n", "No", "Book Id", fmt.Sprintf("%*s", 25, "Title"), "Author", "Status")
							rowLen := len(title)

							fmt.Println("== List of Book ==")
							fmt.Print(title)
							for i := 0; i < rowLen; i++ {
								fmt.Print("=")
							}
							fmt.Print("\n")
							if result != nil {
								i := 1
								for _, value := range result {
									// if value.IDUser != currentUser.ID {
									// 	if !value.Is_Borrowed {
									fmt.Printf("%4d | %s | %25s | %25s | %s |\n", i, fmt.Sprintf("%*d", 7, value.ID), value.Title, value.Author, fmt.Sprintf("%*s", 13, BorrowedStatus(value.Is_Borrowed)))
									i++
									// 	}
									// }
								}
							} else {
								fmt.Println("\n\t\\tt Book Title not Found")
							}

							fmt.Println("\n==============================")
							fmt.Println("1. Choose Book Number to Borrow Book")
							fmt.Println("9. Back")
							fmt.Println("0. Main Menu")
							fmt.Print("Enter Your Input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								var inputBookNumber int
								// validasi input, kalau dia <= 0 panic error out of range
								notValidInput := true
								for notValidInput {
									fmt.Println("== Choose Number to Borrow Book ==")
									fmt.Scanln(&inputBookNumber)
									fmt.Println("Len data resavaild", len(result))
									if inputBookNumber > len(result) || inputBookNumber < 1 {
										Message("Input Not Valid", "Number input must in range.", "")
									} else {
										notValidInput = false
									}
									isFindTitle = false
								}

								callClear()
								var targetedBook model.Books = result[inputBookNumber-1]
								ViewDetailBook(targetedBook)

								// fmt.Println("TEEESS\n\n\n\ntargeted", targetedBook, "currentbook", currentUser, "\n")
								// Validasi tidak bisa meminjam buku milik dia sendiri
								if targetedBook.IDUser == currentUser.ID || targetedBook.Is_Borrowed {
									Message("Borrow Failed", "You can't borrow your own book Or Book is not Available.", "")
								} else {
									fmt.Println("\nAre you sure want to borrow this book? (Y/N)")

									var isNotYesNo bool = true

									for isNotYesNo {
										fmt.Scanln(&yesNo)
										if yesNo == "Y" || yesNo == "y" {
											isNotYesNo = false

											ownerData, err := userControl.GetById(targetedBook.IDUser)
											if err != nil {
												fmt.Println("Failed Retrieve Book's Owner Data.")
											}
											var newLend model.Lends
											newLend.Return_date = time.Time{}
											newLend.Is_returned = false
											newLend.Book_owner_id = ownerData.ID
											newLend.Book_owner_email = ownerData.Email
											newLend.Book_owner_name = ownerData.Name
											newLend.Book_owner_address = ownerData.Address
											newLend.Book_owner_phone = ownerData.Phone_number
											newLend.Book_title = targetedBook.Title
											newLend.Book_author = targetedBook.Author
											newLend.IDUser = currentUser.ID
											newLend.IDBook = targetedBook.ID

											resBorrowBook, err := lendControl.AddLend(newLend)

											if err != nil {
												fmt.Println("Error on Borrow Book", err.Error())
											} else {
												// Update Status Is Borrowed di Buku
												targetedBook.Is_Borrowed = true
												resUpdateIsborrowed, err := bookControl.Edit(targetedBook)
												if err != nil {
													fmt.Println("Error on Update IsBorrowed Status", err.Error())
												} else {
													if resUpdateIsborrowed.ID != 0 {
														Message("Success Borrow Book", "Success Add Borrowed Book : "+resBorrowBook.Book_title, "")
													} else {
														fmt.Println("Error on Update isBorrowedBook, no book updated", err.Error())
													}
												}
											}
											isNotYesNo = false
											isFindTitle = false

										} else if yesNo == "N" || yesNo == "n" {
											isNotYesNo = false
											isFindTitle = false
											callClear()
										} else {
											isNotYesNo = true
										}
									}
								}
							case 9:
								Message("Kembali", "Ke Menu Sebelumnya...", "")
								isFindTitle = false
							case 0:
								Message("Kembali", "Ke Menu...", "")
								isFindTitle = false
								homepageMemberMenu = false
							}
						}
					case 2:
						isFindAuthor := true
						for isFindAuthor {
							callClear()
							fmt.Println("\t--Find Books by Author--")
							fmt.Println("Please enter book's author :")
							fmt.Scanln(&currentBook.Author)

							result, err := bookControl.GetWhere(currentBook.Author, "author")
							if err != nil {
								ErrorMsg(showError, "Searching Book Failed", "Book not found or failed to retrieve data.", err.Error())
							}
							var author string = fmt.Sprintf("%4s | %s | %s | %25s | %13s |\n", "No", "Book Id", fmt.Sprintf("%*s", 25, "Title"), "Author", "Status")
							rowLen := len(author)

							fmt.Println("== List of Book ==")
							fmt.Print(author)
							for i := 0; i < rowLen; i++ {
								fmt.Print("=")
							}
							fmt.Print("\n")
							if result != nil {
								i := 1
								for _, value := range result {
									fmt.Printf("%4d | %s | %25s | %25s | %s |\n", i, fmt.Sprintf("%*d", 7, value.ID), value.Title, value.Author, fmt.Sprintf("%*s", 13, BorrowedStatus(value.Is_Borrowed)))
									i++
								}
							} else {
								fmt.Println("\n\t\\tt Book Title not Found")
							}
							fmt.Println("\n==============================")
							fmt.Println("1. Choose Book Number to Borrow Book")
							fmt.Println("9. Back")
							fmt.Println("0. Main Menu")
							fmt.Print("Enter Your Input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								var inputBookNumber int
								// validasi input, kalau dia <= 0 panic error out of range
								notValidInput := true
								for notValidInput {
									fmt.Println("== Choose Number to Borrow Book ==")
									fmt.Scanln(&inputBookNumber)
									fmt.Println("Len data resavaild", len(result))
									if inputBookNumber > len(result) || inputBookNumber < 1 {
										Message("Input Not Valid", "Number input must in range.", "")
									} else {
										notValidInput = false
									}
									isFindAuthor = false
								}

								callClear()
								var targetedBook model.Books = result[inputBookNumber-1]
								ViewDetailBook(targetedBook)
								// Validasi tidak bisa pinjam buku sendiri
								if targetedBook.IDUser == currentUser.ID || targetedBook.Is_Borrowed {
									Message("Borrow Failed", "You can't borrow your own book Or Book is not Available.", "")
								} else {
									fmt.Println("\nAre you sure want to borrow this book? (Y/N)")

									var isNotYesNo bool = true

									for isNotYesNo {
										fmt.Scanln(&yesNo)
										if yesNo == "Y" || yesNo == "y" {
											isNotYesNo = false

											ownerData, err := userControl.GetById(targetedBook.IDUser)
											if err != nil {
												fmt.Println("Failed Retrieve Book's Owner Data.")
											}
											var newLend model.Lends
											newLend.Return_date = time.Time{}
											newLend.Is_returned = false
											newLend.Book_owner_id = ownerData.ID
											newLend.Book_owner_email = ownerData.Email
											newLend.Book_owner_name = ownerData.Name
											newLend.Book_owner_address = ownerData.Address
											newLend.Book_owner_phone = ownerData.Phone_number
											newLend.Book_title = targetedBook.Title
											newLend.Book_author = targetedBook.Author
											newLend.IDUser = currentUser.ID
											newLend.IDBook = targetedBook.ID

											resBorrowBook, err := lendControl.AddLend(newLend)

											if err != nil {
												fmt.Println("Error on Borrow Book", err.Error())
											} else {
												// Update Status Is Borrowed di Buku
												targetedBook.Is_Borrowed = true
												resUpdateIsborrowed, err := bookControl.Edit(targetedBook)
												if err != nil {
													fmt.Println("Error on Update IsBorrowed Status", err.Error())
												} else {
													if resUpdateIsborrowed.ID != 0 {
														Message("Success Borrow Book", "Success Add Borrowed Book : "+resBorrowBook.Book_title, "")
													} else {
														fmt.Println("Error on Update isBorrowedBook, no book updated", err.Error())
													}
												}
											}
											isNotYesNo = false
											isFindAuthor = false

										} else if yesNo == "N" || yesNo == "n" {
											isNotYesNo = false
											isFindAuthor = false
											callClear()
										} else {
											isNotYesNo = true
										}
									}
								}
							case 9:
								Message("Kembali", "Ke Menu Sebelumnya...", "")
								isFindAuthor = false
							case 0:
								Message("Kembali", "Ke Menu...", "")
								isFindAuthor = false
								homepageMemberMenu = false
							}
						}
					case 9:
						Message("Kembali", "Ke Menu Sebelumnya...", "")
						isRunning = false
					}
				case 2:
					isListAvailableBook := true
					for isListAvailableBook {

						callClear()

						fmt.Println("\t--List Book--")

						resAvailBook, err := bookControl.GetAll()
						if err != nil {
							Message("Failed Get Book", "Failed Retrieve Data", err.Error())
						}

						fmt.Println("List Book Table")
						fmt.Println("==================================")
						fmt.Printf("%4s | %7s | %25s | %25s | %13s |\n", "No", "Book Id", "Title", "Author", "Status")

						if resAvailBook != nil {
							i := 1
							for _, value := range resAvailBook {
								fmt.Printf("%4d | %7d | %25s | %25s | %13s |\n", i, value.ID, value.Title, value.Author, BorrowedStatus(value.Is_Borrowed))
								i++
							}
						} else {
							fmt.Println("\n\t\\tt Book Title not Found")
						}
						fmt.Println("\n==============================")
						fmt.Println("1. Choose Book Number to Borrow Book")
						fmt.Println("9. Back")
						fmt.Println("0. Main Menu")
						fmt.Print("Enter Your Input: ")
						fmt.Scanln(&inputMenu)
						switch inputMenu {
						case 1:
							var inputBookNumber int
							// validasi input, kalau dia <= 0 panic error out of range
							notValidInput := true
							for notValidInput {
								fmt.Println("== Choose Number to Borrow Book ==")
								fmt.Scanln(&inputBookNumber)
								fmt.Println("Len data resavaild", len(resAvailBook))
								if inputBookNumber > len(resAvailBook) || inputBookNumber < 1 {
									Message("Input Not Valid", "Number input must in range.", "")
								} else {
									notValidInput = false
								}
							}

							callClear()
							var targetedBook model.Books = resAvailBook[inputBookNumber-1]
							ViewDetailBook(targetedBook)
							// Validasi tidak bisa pinjam buku sendiri
							if targetedBook.IDUser == currentUser.ID || targetedBook.Is_Borrowed {
								Message("Borrow Failed", "You can't borrow your own book Or Book is not Available.", "")
							} else {
								fmt.Println("\nAre you sure want to borrow this book? (Y/N)")

								var isNotYesNo bool = true

								for isNotYesNo {
									fmt.Scanln(&yesNo)
									if yesNo == "Y" || yesNo == "y" {
										isNotYesNo = false

										ownerData, err := userControl.GetById(targetedBook.IDUser)
										if err != nil {
											fmt.Println("Failed Retrieve Book's Owner Data.")
										}
										var newLend model.Lends
										newLend.Return_date = time.Time{}
										newLend.Is_returned = false
										newLend.Book_owner_id = ownerData.ID
										newLend.Book_owner_email = ownerData.Email
										newLend.Book_owner_name = ownerData.Name
										newLend.Book_owner_address = ownerData.Address
										newLend.Book_owner_phone = ownerData.Phone_number
										newLend.Book_title = targetedBook.Title
										newLend.Book_author = targetedBook.Author
										newLend.IDUser = currentUser.ID
										newLend.IDBook = targetedBook.ID

										resBorrowBook, err := lendControl.AddLend(newLend)

										if err != nil {
											fmt.Println("Error on Borrow Book", err.Error())
										} else {
											// Update Status Is Borrowed di Buku
											targetedBook.Is_Borrowed = true
											resUpdateIsborrowed, err := bookControl.Edit(targetedBook)
											if err != nil {
												fmt.Println("Error on Update IsBorrowed Status", err.Error())
											} else {
												if resUpdateIsborrowed.ID != 0 {
													Message("Success Borrow Book", "Success Add Borrowed Book : "+resBorrowBook.Book_title, "")
												} else {
													fmt.Println("Error on Update isBorrowedBook, no book updated", err.Error())
												}

											}
										}
										isNotYesNo = false

									} else if yesNo == "N" || yesNo == "n" {
										isNotYesNo = false
										callClear()
									} else {
										isNotYesNo = true
									}
								}
							}
						case 9:
							Message("Kembali", "Ke Menu Sebelumnya...", "")
							isListAvailableBook = false
						case 0:
							Message("Kembali", "Ke Menu...", "")
							isListAvailableBook = false
							homepageMemberMenu = false
						}
					}
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
							} else {

								Message("Success", "Updating New Name Success", "Update to "+resUpdateUser.Name)
							}

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
							} else {
								Message("Success", "Updating New Address Success", "Update to "+resUpdateUser.Address)
							}

							updateProfile = false
						case 3:
							callClear()
							fmt.Println("\t--Edit Phone Number--")
							fmt.Println("================")
							fmt.Printf("Phone Number: %s", currentUser.Phone_number)
							fmt.Println("\n================")
							fmt.Println("Input New Phone Number :")
							fmt.Scanln(&currentUser.Phone_number)
							resUpdateUser, err := userControl.Edit(currentUser)
							if err != nil {
								ErrorMsg(showError, "Error on Updating Phone Number", "Please try again.", err.Error())
							} else {
								Message("Success", "Updating New Phone Number Success", "Update to "+resUpdateUser.Phone_number)
							}

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
										} else {
											Message("Success", "Updating New Password", "Please "+resUpdateUser.Email+" relogin using new password.")
											isLoggedIn = false
											currentUser = model.Users{}
										}
										isChangePassword = false
										updateProfile = false
										callClear()
									} else {
										Message("Failed", "New Password is not match", "Please try again")
										callClear()
									}
								} else {
									Message("Failed", "Old Password is not correct", "Please try again")
									isChangePassword = false
									callClear()
								}
							}
						case 5:
							callClear()
							statusAccount := "Active"
							if currentUser.Is_Active {
								statusAccount = "Active"
							} else {
								statusAccount = "Not Active"
							}

							fmt.Println("\t--Status Account--")
							fmt.Println("================")
							fmt.Printf("Status: %s", statusAccount)
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
									Message("Success", "Deactivate Account Success", "Your Account "+resUpdateUser.Email+" Now Deactive.")
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
							Message("Kembali", "Ke Menu Sebelumnya...", "")
							updateProfile = false
						case 0:
							Message("Kembali", "Ke Menu Sebelumnya...", "")
							updateProfile = false
							homepageMemberMenu = false
						}
					}
				case 4:
					callClear()
					myLibraryMenu := true
					for myLibraryMenu {
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

									result, err := bookControl.GetUserBooks(currentUser.ID)
									if err != nil {
										ErrorMsg(showError, "Failed Retrieve List Books", "Please try again.", err.Error())
									}

									fmt.Println("List My Book Table")
									fmt.Println("==================================")
									fmt.Printf("%4s | %7s | %25s | %25s | %13s |\n", "No", "Book Id", "Title", "Author", "Status")

									if !reflect.ValueOf(result).IsZero() {
										i := 1
										for _, value := range result {
											fmt.Printf("%4d | %7d | %25s | %25s | %13s |\n", i, value.ID, value.Title, value.Author, BorrowedStatus(value.Is_Borrowed))
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
										// validasi input, kalau dia <= 0 panic error out of range
										notValidInput := true
										for notValidInput {
											fmt.Println("== Choose Number Edit Book Data ==")
											fmt.Scanln(&inputBookNumber)

											if inputBookNumber > len(result) || inputBookNumber < 1 {
												Message("Input Not Valid", "Number input must in range.", "")
											} else {
												notValidInput = false
											}
										}

										isEditBookMenu := true

										for isEditBookMenu {

											result, _ := bookControl.GetUserBooks(currentUser.ID)
											var targetedBook model.Books = result[inputBookNumber-1]

											callClear()
											fmt.Println("Edit My Book")
											ViewDetailBook(targetedBook)
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
												} else {
													Message("Success", "Edit Title Success", "Update to "+resUpdateBook.Title)
												}

											case 2:
												fmt.Println("\t--Edit Author--")
												fmt.Println("Input Author :")
												fmt.Scanln(&targetedBook.Author)
												resUpdateBook, err := bookControl.Edit(targetedBook)
												if err != nil {
													ErrorMsg(showError, "Error on Updating Book", "Please try again", err.Error())
												} else {
													Message("Success", "Edit Author Success", "Update to"+resUpdateBook.Author)
												}
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
														} else {
															Message("Success", "Deleting Book Success", "Delete Data"+resultDelete.Title)
														}

														isNotYesNo = false
														callClear()
													} else if yesNo == "N" || yesNo == "n" {
														isNotYesNo = false
														callClear()
													} else {
														isNotYesNo = true
													}
												}
												isEditBookMenu = false
											case 9:
												isEditBookMenu = false
												listMyBookMenu = false
											case 0:
												isEditBookMenu = false
												listMyBookMenu = false

											}
											result, _ = bookControl.GetUserBooks(currentUser.ID)
										}

									case 9:
										Message("Kembali", "Ke Menu Sebelumnya...", "")
										listMyBookMenu = false
									case 0:
										Message("Kembali", "Ke Menu...", "")
										listMyBookMenu = false
										myLibraryMenu = false
										homepageMemberMenu = false
									}
								}
							}
						case 2:
							Borrow := true
							for Borrow {
								callClear()
								fmt.Println("\t--Borrow--")
								fmt.Println("1. Search Available Book")
								fmt.Println("2. List of Available Books")
								fmt.Println("3. List of Borrowed Books")
								fmt.Println("===================")
								fmt.Println("9. Back")
								fmt.Println("0. Main Menu")
								fmt.Print("Enter Your Input: ")
								fmt.Scanln(&inputMenu)

								switch inputMenu {
								case 1:
									callClear()
									fmt.Println("\t--Search Available Books to Borrow--")
									fmt.Println("1. Find Books by Title")
									fmt.Println("2. Find Books by Author")
									fmt.Println("9. Back")
									fmt.Print("Enter Your Input: ")
									fmt.Scanln(&inputMenu)
									switch inputMenu {
									case 1:
										isFindTitle := true

										for isFindTitle {
											callClear()
											fmt.Println("\t--Find Books by Title--")
											fmt.Println("Please enter book's title :")
											fmt.Scanln(&currentBook.Title)

											result, err := bookControl.GetWhere(currentBook.Title, "title")
											if err != nil {
												ErrorMsg(showError, "Searching Book Failed", "Book not found or failed to retrieve data.", err.Error())
											}
											var title string = fmt.Sprintf("%4s | %s | %s | %25s | %13s |\n", "No", "Book Id", fmt.Sprintf("%*s", 25, "Title"), "Author", "Status")
											rowLen := len(title)

											fmt.Println("== List of Book ==")
											fmt.Print(title)
											for i := 0; i < rowLen; i++ {
												fmt.Print("=")
											}
											fmt.Print("\n")
											if result != nil {
												i := 1
												for _, value := range result {
													if value.IDUser != currentUser.ID {
														if !value.Is_Borrowed {
															fmt.Printf("%4d | %s | %25s | %25s | %s |\n", i, fmt.Sprintf("%*d", 7, value.ID), value.Title, value.Author, fmt.Sprintf("%*s", 13, BorrowedStatus(value.Is_Borrowed)))
															i++
														}
													}
												}
											} else {
												fmt.Println("\n\t\\tt Book Title not Found")
											}

											fmt.Println("\n==============================")
											fmt.Println("1. Choose Book Number to Borrow Book")
											fmt.Println("9. Back")
											fmt.Println("0. Main Menu")
											fmt.Print("Enter Your Input: ")
											fmt.Scanln(&inputMenu)
											switch inputMenu {
											case 1:
												var inputBookNumber int
												// validasi input, kalau dia <= 0 panic error out of range
												notValidInput := true
												for notValidInput {
													fmt.Println("== Choose Number to Borrow Book ==")
													fmt.Scanln(&inputBookNumber)
													fmt.Println("Len data resavaild", len(result))
													if inputBookNumber > len(result) || inputBookNumber < 1 {
														Message("Input Not Valid", "Number input must in range.", "")
													} else {
														notValidInput = false
													}
													isFindTitle = false
												}

												callClear()
												var targetedBook model.Books = result[inputBookNumber-1]
												ViewDetailBook(targetedBook)

												// Validasi tidak bisa meminjam buku milik dia sendiri
												if targetedBook.IDUser == currentUser.ID || targetedBook.Is_Borrowed {
													Message("Borrow Failed", "You can't borrow your own book Or Book is not Available.", "")
												} else {
													fmt.Println("\nAre you sure want to borrow this book? (Y/N)")

													var isNotYesNo bool = true

													for isNotYesNo {
														fmt.Scanln(&yesNo)
														if yesNo == "Y" || yesNo == "y" {
															isNotYesNo = false

															ownerData, err := userControl.GetById(targetedBook.IDUser)
															if err != nil {
																fmt.Println("Failed Retrieve Book's Owner Data.")
															}
															var newLend model.Lends
															newLend.Return_date = time.Time{}
															newLend.Is_returned = false
															newLend.Book_owner_id = ownerData.ID
															newLend.Book_owner_email = ownerData.Email
															newLend.Book_owner_name = ownerData.Name
															newLend.Book_owner_address = ownerData.Address
															newLend.Book_owner_phone = ownerData.Phone_number
															newLend.Book_title = targetedBook.Title
															newLend.Book_author = targetedBook.Author
															newLend.IDUser = currentUser.ID
															newLend.IDBook = targetedBook.ID

															resBorrowBook, err := lendControl.AddLend(newLend)

															if err != nil {
																fmt.Println("Error on Borrow Book", err.Error())
															} else {
																// Update Status Is Borrowed di Buku
																targetedBook.Is_Borrowed = true
																resUpdateIsborrowed, err := bookControl.Edit(targetedBook)
																if err != nil {
																	fmt.Println("Error on Update IsBorrowed Status", err.Error())
																} else {
																	if resUpdateIsborrowed.ID != 0 {
																		Message("Success Borrow Book", "Success Add Borrowed Book : "+resBorrowBook.Book_title, "")
																	} else {
																		fmt.Println("Error on Update isBorrowedBook, no book updated", err.Error())
																	}
																}
															}
															isNotYesNo = false
															isFindTitle = false

														} else if yesNo == "N" || yesNo == "n" {
															isNotYesNo = false
															isFindTitle = false
															callClear()
														} else {
															isNotYesNo = true
														}
													}
												}
											case 9:
												Message("Kembali", "Ke Menu Sebelumnya...", "")
												isFindTitle = false
											case 0:
												Message("Kembali", "Ke Menu...", "")
												isFindTitle = false
												homepageMemberMenu = false
											}
										}
									case 2:
										isFindAuthor := true
										for isFindAuthor {
											callClear()
											fmt.Println("\t--Find Books by Author--")
											fmt.Println("Please enter book's author :")
											fmt.Scanln(&currentBook.Author)

											result, err := bookControl.GetWhere(currentBook.Author, "author")
											if err != nil {
												ErrorMsg(showError, "Searching Book Failed", "Book not found or failed to retrieve data.", err.Error())
											}
											var author string = fmt.Sprintf("%4s | %s | %s | %25s | %13s |\n", "No", "Book Id", fmt.Sprintf("%*s", 25, "Title"), "Author", "Status")
											rowLen := len(author)

											fmt.Println("== List of Book ==")
											fmt.Print(author)
											for i := 0; i < rowLen; i++ {
												fmt.Print("=")
											}
											fmt.Print("\n")
											if result != nil {
												i := 1
												for _, value := range result {
													// validasi
													if value.IDUser != currentUser.ID {
														if !value.Is_Borrowed {
															fmt.Printf("%4d | %s | %25s | %25s | %s |\n", i, fmt.Sprintf("%*d", 7, value.ID), value.Title, value.Author, fmt.Sprintf("%*s", 13, BorrowedStatus(value.Is_Borrowed)))
															i++
														}
													}
												}
											} else {
												fmt.Println("\n\t\\tt Book Title not Found")
											}
											fmt.Println("\n==============================")
											fmt.Println("1. Choose Book Number to Borrow Book")
											fmt.Println("9. Back")
											fmt.Println("0. Main Menu")
											fmt.Print("Enter Your Input: ")
											fmt.Scanln(&inputMenu)
											switch inputMenu {
											case 1:
												var inputBookNumber int
												// validasi input, kalau dia <= 0 panic error out of range
												notValidInput := true
												for notValidInput {
													fmt.Println("== Choose Number to Borrow Book ==")
													fmt.Scanln(&inputBookNumber)
													fmt.Println("Len data resavaild", len(result))
													if inputBookNumber > len(result) || inputBookNumber < 1 {
														Message("Input Not Valid", "Number input must in range.", "")
													} else {
														notValidInput = false
													}
													isFindAuthor = false
												}

												callClear()
												var targetedBook model.Books = result[inputBookNumber-1]
												ViewDetailBook(targetedBook)
												// Validasi tidak bisa pinjam buku sendiri
												if targetedBook.IDUser == currentUser.ID || targetedBook.Is_Borrowed {
													Message("Borrow Failed", "You can't borrow your own book Or Book is not Available.", "")
												} else {
													fmt.Println("\nAre you sure want to borrow this book? (Y/N)")

													var isNotYesNo bool = true

													for isNotYesNo {
														fmt.Scanln(&yesNo)
														if yesNo == "Y" || yesNo == "y" {
															isNotYesNo = false

															ownerData, err := userControl.GetById(targetedBook.IDUser)
															if err != nil {
																fmt.Println("Failed Retrieve Book's Owner Data.")
															}
															var newLend model.Lends
															newLend.Return_date = time.Time{}
															newLend.Is_returned = false
															newLend.Book_owner_id = ownerData.ID
															newLend.Book_owner_email = ownerData.Email
															newLend.Book_owner_name = ownerData.Name
															newLend.Book_owner_address = ownerData.Address
															newLend.Book_owner_phone = ownerData.Phone_number
															newLend.Book_title = targetedBook.Title
															newLend.Book_author = targetedBook.Author
															newLend.IDUser = currentUser.ID
															newLend.IDBook = targetedBook.ID

															resBorrowBook, err := lendControl.AddLend(newLend)

															if err != nil {
																fmt.Println("Error on Borrow Book", err.Error())
															} else {
																// Update Status Is Borrowed di Buku
																targetedBook.Is_Borrowed = true
																resUpdateIsborrowed, err := bookControl.Edit(targetedBook)
																if err != nil {
																	fmt.Println("Error on Update IsBorrowed Status", err.Error())
																} else {
																	if resUpdateIsborrowed.ID != 0 {
																		Message("Success Borrow Book", "Success Add Borrowed Book : "+resBorrowBook.Book_title, "")
																	} else {
																		fmt.Println("Error on Update isBorrowedBook, no book updated", err.Error())
																	}
																}
															}
															isNotYesNo = false
															isFindAuthor = false

														} else if yesNo == "N" || yesNo == "n" {
															isNotYesNo = false
															isFindAuthor = false
															callClear()
														} else {
															isNotYesNo = true
														}
													}
												}
											case 9:
												Message("Kembali", "Ke Menu Sebelumnya...", "")
												isFindAuthor = false
											case 0:
												Message("Kembali", "Ke Menu...", "")
												isFindAuthor = false
												homepageMemberMenu = false
											}
										}
									case 9:
										Message("Kembali", "Ke Menu Sebelumnya...", "")
										Borrow = false
									}

								case 2:
									isListAvailableBook := true
									for isListAvailableBook {

										callClear()

										fmt.Println("\t--List Available Book to Borrow--")

										resAvailBook, err := bookControl.GetAllAvailable(currentUser.ID)
										if err != nil {
											Message("Failed Get Avaialble Book", "Failed Retrieve Data", err.Error())
										}

										fmt.Println("List Available Book Table")
										fmt.Println("==================================")
										fmt.Printf("%4s | %7s | %25s | %25s | %13s |\n", "No", "Book Id", "Title", "Author", "Status")
										if resAvailBook != nil {
											i := 1
											for _, value := range resAvailBook {
												// validasi
												if value.IDUser != currentUser.ID {
													if !value.Is_Borrowed {
														fmt.Printf("%4d | %7d | %25s | %25s | %13s |\n", i, value.ID, value.Title, value.Author, BorrowedStatus(value.Is_Borrowed))
														i++
													}
												}
											}
										} else {
											fmt.Println("\n\t\\tt Book Title not Found")
										}
										fmt.Println("\n==============================")
										fmt.Println("1. Choose Book Number to Borrow Book")
										fmt.Println("9. Back")
										fmt.Println("0. Main Menu")
										fmt.Print("Enter Your Input: ")
										fmt.Scanln(&inputMenu)
										switch inputMenu {
										case 1:
											var inputBookNumber int
											// validasi input, kalau dia <= 0 panic error out of range
											notValidInput := true
											for notValidInput {
												fmt.Println("== Choose Number to Borrow Book ==")
												fmt.Scanln(&inputBookNumber)
												fmt.Println("Len data resavaild", len(resAvailBook))
												if inputBookNumber > len(resAvailBook) || inputBookNumber < 1 {
													Message("Input Not Valid", "Number input must in range.", "")
												} else {
													notValidInput = false
												}
											}

											callClear()
											var targetedBook model.Books = resAvailBook[inputBookNumber-1]
											ViewDetailBook(targetedBook)
											// Validasi
											if targetedBook.IDUser == currentUser.ID || targetedBook.Is_Borrowed {
												Message("Borrow Failed", "You can't borrow your own book Or Book is not Available.", "")
											} else {
												fmt.Println("\nAre you sure want to borrow this book? (Y/N)")

												var isNotYesNo bool = true

												for isNotYesNo {
													fmt.Scanln(&yesNo)
													if yesNo == "Y" || yesNo == "y" {
														isNotYesNo = false

														ownerData, err := userControl.GetById(targetedBook.IDUser)
														if err != nil {
															fmt.Println("Failed Retrieve Book's Owner Data.")
														}
														var newLend model.Lends
														newLend.Return_date = time.Time{}
														newLend.Is_returned = false
														newLend.Book_owner_id = ownerData.ID
														newLend.Book_owner_email = ownerData.Email
														newLend.Book_owner_name = ownerData.Name
														newLend.Book_owner_address = ownerData.Address
														newLend.Book_owner_phone = ownerData.Phone_number
														newLend.Book_title = targetedBook.Title
														newLend.Book_author = targetedBook.Author
														newLend.IDUser = currentUser.ID
														newLend.IDBook = targetedBook.ID

														resBorrowBook, err := lendControl.AddLend(newLend)

														if err != nil {
															fmt.Println("Error on Borrow Book", err.Error())
														} else {
															// Update Status Is Borrowed di Buku
															targetedBook.Is_Borrowed = true
															resUpdateIsborrowed, err := bookControl.Edit(targetedBook)
															if err != nil {
																fmt.Println("Error on Update IsBorrowed Status", err.Error())
															} else {
																if resUpdateIsborrowed.ID != 0 {
																	Message("Success Borrow Book", "Success Add Borrowed Book : "+resBorrowBook.Book_title, "")
																} else {
																	fmt.Println("Error on Update isBorrowedBook, no book updated", err.Error())
																}

															}
														}
														isNotYesNo = false

													} else if yesNo == "N" || yesNo == "n" {
														isNotYesNo = false
														callClear()
													} else {
														isNotYesNo = true
													}
												}
											}
										case 9:
											Message("Kembali", "Ke Menu Sebelumnya...", "")
											isListAvailableBook = false
										case 0:
											Message("Kembali", "Ke Menu...", "")
											isListAvailableBook = false
											Borrow = false
											myLibraryMenu = false
											homepageMemberMenu = false
										}
									}
								case 3:
									listBorrowedMenu := true
									for listBorrowedMenu {

										callClear()
										var inputBookNumber int

										result, err := lendControl.GetUserBookBorrow(currentUser.ID)
										if err != nil {
											Message("Get Borrowed List Failed", "Failed to retieve Borrowed List", err.Error())
										}

										fmt.Println("List of My Borrowed Books")
										fmt.Println("==================================")
										fmt.Printf("%4s | %7s | %25s | %25s | %10s |\n", "No", "Book Id", "Title", "Author", "Book Owner")

										if result != nil {
											i := 1
											for _, value := range result {
												fmt.Printf("%4d | %7d | %25s | %25s | %10s |\n", i, value.ID, value.Book_title, value.Book_author, value.Book_owner_name)
												i++
											}
										} else {
											fmt.Println("\n\t\\tt Book Title not Found")
										}

										fmt.Println("\n==============================")
										fmt.Println("1. Choose the book to be returned")
										fmt.Println("9. Back")
										fmt.Println("0. Main Menu")
										fmt.Print("Enter Your Input: ")
										fmt.Scanln(&inputMenu)
										switch inputMenu {
										case 1:
											// validasi input, kalau dia <= 0 panic error out of range
											notValidInput := true
											for notValidInput {
												fmt.Println("== Choose Number To Be Returned ==")
												fmt.Scanln(&inputBookNumber)

												if inputBookNumber > len(result) || inputBookNumber < 1 {
													Message("Input Not Valid", "Number input must in range.", "")
												} else {
													notValidInput = false
												}
											}

											var targetedLend model.Lends = result[inputBookNumber-1]

											callClear()
											fmt.Println("List My Borrowed Books")
											fmt.Println("==================================")
											fmt.Printf("%4s | %7s | %25s | %25s | %10s |\n", "No", "Book Id", "Title", "Author", "Book Owner")

											fmt.Printf("%4d | %7d | %25s | %25s | %10s |\n", 1, targetedLend.ID, targetedLend.Book_title, targetedLend.Book_author, targetedLend.Book_owner_name)

											fmt.Println("\n==============================")

											isNotYesNo := true
											for isNotYesNo {
												fmt.Println("Are you sure to Return this book? (y/n)")
												fmt.Scanln(&yesNo)
												if yesNo == "Y" || yesNo == "y" {
													targetedLend.Is_returned = true
													resultReturn, err := lendControl.Model.ReturnBook(targetedLend)
													if err != nil {
														Message("Failed", "Returning Book Failed", resultReturn)
														fmt.Println("", err.Error())
													} else {
														// Update Status Is Borrowed di Buku menjadi false = available
														targetedBook, err := bookControl.GetById(resultReturn.IDBook)

														if err != nil {
															fmt.Println("Failed get book to update status")
														} else {
															targetedBook.Is_Borrowed = false
															resUpdateIsborrowed, err := bookControl.Edit(targetedBook)
															if err != nil {
																fmt.Println("Error on Update IsBorrowed Status to Available", err.Error())
															} else {
																if resUpdateIsborrowed.ID > 0 {
																	Message("Success Return Book", "Success Returning Book : "+resultReturn.Book_title, "")
																} else {
																	fmt.Println("Error on Update isBorrowedBook to Available, no book updated", err.Error())
																}

															}
														}
													}

													isNotYesNo = false
													Borrow = false
													callClear()
												} else if yesNo == "N" || yesNo == "n" {
													isNotYesNo = false
													Borrow = false
													callClear()
												} else {
													isNotYesNo = true
												}
											}
										case 9:
											Message("Kembali", "Ke Menu Sebelumnya...", "")
											listBorrowedMenu = false
										case 0:
											Message("Kembali", "Ke Menu...", "")
											listBorrowedMenu = false
											homepageMemberMenu = false
										}
									}
								case 9:
									Message("Kembali", "Ke Menu Sebelumnya...", "")
									Borrow = false
								case 0:
									Message("Kembali", "Ke Menu Sebelumnya...", "")
									Borrow = false
									homepageMemberMenu = false
								}
							}
						case 9:
							Message("Kembali", "Ke Menu Sebelumnya...", "")
							myLibraryMenu = false
						case 0:
							Message("Kembali", "Ke Menu Sebelumnya...", "")
							myLibraryMenu = false
							homepageMemberMenu = false
						}
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
}

func Message(_title, _content, _detail interface{}) {
	var next string
	fmt.Println("\n=== ", _title, " ===")
	fmt.Println(_content)
	fmt.Println(_detail)
	fmt.Println("Please Enter to continue...")
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
	fmt.Println("\nPlease Enter to continue...")
	fmt.Scanln(&next)
}

func ViewDetailBook(_bookData model.Books) {

	var bookStatus string
	if _bookData.Is_Borrowed {
		bookStatus = "Not Available"
	} else {
		bookStatus = "Available"
	}

	fmt.Println("===== Book Detail =====")
	fmt.Println("Book Title : ", _bookData.Title)
	fmt.Println("Book Author : ", _bookData.Author)
	fmt.Println("Book Status : ", bookStatus)
}

func BorrowedStatus(Is_Borrowed bool) string {
	if Is_Borrowed {
		return "Not Available"
	} else {
		return "Available"
	}
}
