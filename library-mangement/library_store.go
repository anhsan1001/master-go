package librarymangement

import (
	"fmt"
	"main/models"
	"strings"
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

func (lib *Library) ListBorrowHistoryByBorrowerID(borrowerID string) []models.Transaction {

	if _, borrowerExits := lib.borrower[borrowerID]; !borrowerExits {
		return nil
	}
	var history []models.Transaction
	for _, transaction := range lib.transaction {
		if transaction.BorrowerID == borrowerID {
			history = append(history, transaction)
		}
	}

	return history
}

func (lib *Library) ReturnBookStore(transactionID string) error {
	transaction, exitsTransaction := lib.transaction[transactionID]
	if !exitsTransaction {
		return fmt.Errorf("transaction khong ton tai")
	}
	book, bookExits := lib.book[transaction.BookID]
	if !bookExits {
		return fmt.Errorf("sach khong ton tai")
	}
	book.Status = "available"
	lib.book[transaction.BookID] = book
	transaction.ReturnDate = time.Now()
	lib.transaction[transaction.ID] = transaction
	return nil
}

func (lib *Library) SearchBookStore(queryString string) []models.Book {
	var queryToLower = strings.ToLower(queryString)
	var result []models.Book
	for _, book := range lib.book {
		if strings.Contains(strings.ToLower(book.Title), queryToLower) || strings.Contains(strings.ToLower(book.Author), queryToLower) {
			result = append(result, book)
		}
	}
	return result
}
