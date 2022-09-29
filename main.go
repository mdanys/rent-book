package main

import (
	"fmt"
	"os"
	"os/exec"
	"rent-book/controllers"
	"rent-book/models"
	model "rent-book/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	db.AutoMigrate(&model.Lends{})
}

func main() {
	var inputMenu int
	var next string
	gconn, err := connectGorm()
	migrate(gconn)
	userModel := models.UsersModel{DB: gconn}
	userControl := controllers.UsersControl{Model: userModel}
	bookModel := models.BooksModel{DB: gconn}
	bookControl := controllers.BooksControl{Model: bookModel}
	lendModel := models.LendsModel{DB: gconn}
	lendControl := controllers.LendsControl{Model: lendModel}

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
						//Message("Pencarian Buku Gagal", "Buku tidak di temukan", err.Error())
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
							fmt.Println("Error on Adding Book", err.Error())
						}
						// fmt.Println("Input :", newUsers)
						fmt.Println("\nAdding Book Success", result)
					case 2:
						listMyBookMenu := true
						for listMyBookMenu {

							callClear()
							var inputBookNumber int

							fmt.Println("\t--List My Book--")

							result, err := bookControl.GetAll()
							if err != nil {
								//Message("Pencarian Buku Gagal", "Buku tidak di temukan", err.Error())
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
										fmt.Println("Error on Updating Book", err.Error())
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
										fmt.Println("Error on Updating Book", err.Error())
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
												Message("Failed", "Deleting Book Failed", resultDelete)
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
					Borrow := true
					for Borrow {
						callClear()
						fmt.Println("\t--Borrow--")
						fmt.Println("1. Search Book")
						fmt.Println("2. List of Books")
						fmt.Println("3. List of Borrowed Books")
						fmt.Println("===================")
						fmt.Println("9. Back")
						fmt.Println("0. Main Menu")
						fmt.Print("Enter Your Input: ")
						fmt.Scanln(&inputMenu)

						switch inputMenu {
						case 1:
						case 2:
							isListAvailableBook := true
							for isListAvailableBook {

								callClear()

								fmt.Println("\t--List Available Book--")

								resAvailBook, err := bookControl.GetAllAvailable()
								if err != nil {
									Message("Failed Get Avaialble Book", "Failed Retrieve Data", err.Error())
								}

								fmt.Println("List Available Book Table")
								fmt.Println("==================================")
								fmt.Printf("%4s | %5s | %15s | %15s | %15s | %15s |\n", "No", "Book Id", "Title", "Author", "Status", "Owner")

								if resAvailBook != nil {
									i := 1
									var status string
									for _, value := range resAvailBook {
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
								fmt.Println("1. Choose Book Number to Borrow Book")
								fmt.Println("9. Back")
								fmt.Println("0. Main Menu")
								fmt.Print("Enter Your Input: ")
								fmt.Scanln(&inputMenu)
								switch inputMenu {
								case 1:
									var inputBookNumber int
									fmt.Println("== Choose Number to Borrow Book ==")
									fmt.Scanln(&inputBookNumber)

									callClear()
									var targetedBook model.Books = resAvailBook[inputBookNumber-1]
									ViewDetailBook(targetedBook)

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
								case 9:
									isListAvailableBook = false
								case 0:

								}
							}
						case 3:
							callClear()
							var inputBookNumber int

							result, err := lendControl.GetUserBookBorrow(currentUser.ID)
							if err != nil {
								Message("Get Borrowed List Failed", "Failed to retieve Borrowed List", err.Error())
							}

							fmt.Println("List of My Borrowed Books")
							fmt.Println("==================================")
							fmt.Printf("%4s | %5s | %15s | %15s | %15s |\n", "No", "Book Id", "Title", "Author", "Book Owner")

							if result != nil {
								i := 1
								for _, value := range result {
									fmt.Printf("%4d | %5d | %15s | %15s | %15s |\n", i, value.ID, value.Book_title, value.Book_author, value.Book_owner_name)
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
								fmt.Println("== Choose Number To Be Returned ==")
								fmt.Scanln(&inputBookNumber)

								var targetedLend model.Lends = result[inputBookNumber-1]

								callClear()
								fmt.Println("List My Borrowed Books")
								fmt.Println("==================================")
								fmt.Printf("%4s | %5s | %15s | %15s | %15s |\n", "No", "Book Id", "Title", "Author", "Book Owner")

								fmt.Printf("%4d | %5d | %15s | %15s | %15s |\n", 1, targetedLend.ID, targetedLend.Book_title, targetedLend.Book_author, targetedLend.Book_owner_name)

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
										Message("Success", "Returning Book Success", resultReturn)

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
							case 0:
							}
						case 9:
						case 0:

						}
					}
				case 9:
				case 0:
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

func ViewDetailBook(_bookData model.Books) {

	var bookStatus string
	if _bookData.Is_Borrowed {
		bookStatus = "Not Available"
	} else {
		bookStatus = "Available"
	}

	fmt.Println("======Book Detail======")
	fmt.Println("Book Title : ", _bookData.Title)
	fmt.Println("Book Author : ", _bookData.Author)
	fmt.Println("Book Status : ", bookStatus)
}
