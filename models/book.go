package models

type Book struct {
	ID     string
	Title  string
	Author string
	Status string
}

func New(b Book) *Book {

	return &Book{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Status: b.Status,
	}
}
