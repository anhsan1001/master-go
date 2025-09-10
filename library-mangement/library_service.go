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
func ListBorrowerHistory(lib *Library) error {
	borrowerID := utils.GetNonEmptyString("Nhap ID nguoi muon: ")
	transactions := lib.ListBorrowHistoryByBorrowerID(borrowerID)
	if len(transactions) == 0 {
		fmt.Println("Khong co lich su muon")
		return nil
	}
	fmt.Println("Lich su muon sach")
	for _, transaction := range transactions {
		fmt.Printf("%s | %s | %s| %s\n", transaction.ID, transaction.BookID, transaction.BorrowerID, transaction.ReturnDate.Format("2006-01-01"))
	}
	return nil

}

func ReturnBook(lib *Library) error {

	transationID := utils.GetNonEmptyString("Nhap ID transation: ")
	lib.ReturnBookStore(transationID)
	return nil

}
func SearchBook(lib *Library) error {
	query := utils.GetNonEmptyString("Nhap ten hoac tac gia de tim kiem sach: ")
	books := lib.SearchBookStore(query)
	if len(books) == 0 {
		fmt.Println("Sach nay khong ton tai")
	}
	for _, book := range books {

		fmt.Printf("ID: %v | Title: %v | Author: %v | Status: %v\n", book.ID, book.Title, book.Author, book.Status)
	}
	return nil
}
