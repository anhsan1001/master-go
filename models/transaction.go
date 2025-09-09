package models

import "time"

type Transaction struct {
	ID           string
	BookID       string
	BorrowerID   string
	BorrowerDate time.Time
	ReturnDate   time.Time
}
