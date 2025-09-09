package librarymangement

import (
	"fmt"
	"main/models"
	"time"
)

type Library struct {
	book        map[string]models.Book
	borrower    map[string]models.Borrower
	transaction map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		book:        make(map[string]models.Book),
		borrower:    make(map[string]models.Borrower),
		transaction: make(map[string]models.Transaction),
	}
}

func (lib *Library) AddBookStore(id, title, author string) {

	lib.book[id] = models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "available",
	}
}
func (lib *Library) BookStoreList() []models.Book {
	books := make([]models.Book, 0, len(lib.book))
	for _, book := range lib.book {
		books = append(books, book)
	}
	return books
}
func (lib *Library) AddBorrowerStore(id, name, email string) {

	lib.borrower[id] = models.Borrower{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func (lib *Library) BorrowerStoreList() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrower))
	for _, borrower := range lib.borrower {
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func (lib *Library) AddTransactionStore(id, bookID, borrowerID string) error {

	book, bookExist := lib.book[bookID]
	if !bookExist {
		return fmt.Errorf("sach khong ton tai")
	}

	if book.Status == "da_muon" {
		return fmt.Errorf("sach da duoc muon")
	}
	if _, borrowerExits := lib.borrower[borrowerID]; !borrowerExits {
		return fmt.Errorf("nguoi muon khong ton tai")
	}

	book.Status = "da_muon"
	lib.book[bookID] = book
	lib.transaction[id] = models.Transaction{
		ID:           id,
		BookID:       bookID,
		BorrowerID:   borrowerID,
		BorrowerDate: time.Now(),
	}
	return nil
}
