package librarymangement

import (
	"fmt"
	"main/utils"
)

func AddBook(lib *Library) error {
	id := utils.GenerateID()
	title := utils.GetNonEmptyString("Nhap tieu de: ")
	author := utils.GetNonEmptyString("Nhap tac gia: ")
	lib.AddBookStore(id, title, author)

	fmt.Println("Them sach thanh cong !")
	return nil
}
func ListBook(lib *Library) error {
	fmt.Println("-------------Danh Sach Sach-------------")
	for _, book := range lib.BookStoreList() {
		fmt.Printf("ID: %v | Title: %v | Author: %v | Status: %v\n", book.ID, book.Title, book.Author, book.Status)
	}
	return nil

}
func AddBorrower(lib *Library) error {
	id := utils.GenerateID()
	name := utils.GetNonEmptyString("Nhap ten: ")
	email := utils.GetNonEmptyString("Nhap email: ")
	lib.AddBorrowerStore(id, name, email)

	fmt.Println("Them nguoi muon thanh cong !")
	return nil

}
func ListBorrower(lib *Library) error {
	fmt.Println("-------------Danh Sach Sach-------------")
	for _, borrower := range lib.BorrowerStoreList() {
		fmt.Printf("ID: %v | Name: %v | Email: %v \n", borrower.ID, borrower.Name, borrower.Email)
	}
	return nil

}
func BorrowerBook(lib *Library) error {
	id := utils.GenerateID()
	bookID := utils.GetNonEmptyString("Nhap book ID: ")
	borrowerID := utils.GetNonEmptyString("Nhap borrower ID: ")

	if err := lib.AddTransactionStore(id, bookID, borrowerID); err != nil {
		return err
	}

	fmt.Println("Muon sach thanh cong !")
	return nil

}
func ListBorrowerHistory() error {
	return nil

}

func ReturnBook() error {

	return nil

}
func SearchBook() error {
	return nil

}
