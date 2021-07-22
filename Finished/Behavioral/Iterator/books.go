package main

import "fmt"

// BookType represents the type of book
type BookType int

// Predefined Book types
const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

// Book represents data about a book
type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}

// Library holds the collection of books
type Library struct {
	Collection []Book
}

// IterateBooks calls the given callback function 
// for each book in the collection
func (l *Library) IterateBooks(f func(Book) error) {
	var err error
	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println("Error encountered:", err)
			break
		}
	}
}

// createIterator returns a BookIterator that can access the book
// collection on demand
func (l *Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}

// -------------------
// Create a Library structure to hold a set of Books
var lib *Library = &Library{
	Collection: []Book{
		{
			Name:      "War and Peace",
			Author:    "Leo Tolstoy",
			PageCount: 864,
			Type:      HardCover,
		},
		{
			Name:      "Crime and Punishment",
			Author:    "Leo Tolstoy",
			PageCount: 1225,
			Type:      SoftCover,
		},
		{
			Name:      "Brave New World",
			Author:    "Aldous Huxley",
			PageCount: 325,
			Type:      PaperBack,
		},
		{
			Name:      "Catcher in the Rye",
			Author:    "J.D. Salinger",
			PageCount: 206,
			Type:      HardCover,
		},
		{
			Name:      "To Kill a Mockingbird",
			Author:    "Harper Lee",
			PageCount: 399,
			Type:      PaperBack,
		},
		{
			Name:      "The Grapes of Wrath",
			Author:    "John Steinbeck",
			PageCount: 464,
			Type:      HardCover,
		},
		{
			Name:      "Wuthering Heights",
			Author:    "Emily Bronte",
			PageCount: 288,
			Type:      EBook,
		},
	},
}
