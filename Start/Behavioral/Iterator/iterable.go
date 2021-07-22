package main

// The IterableCollection interface defines the createIterator
// function, which returns an iterator object
type IterableCollection interface {
	createIterator() iterator
}

// The iterator interface contains the hasNext and next functions
// which allow the collection to return items as needed
type iterator interface {
	hasNext() bool
	next() *Book
}

// TODO: BookIterator is a concrete iterator for a Book collection
type BookIterator struct {
}

func (b *BookIterator) hasNext() bool {
	// TODO: implement hasNext()
	return false
}

func (b *BookIterator) next() *Book {
	// TODO: implement next()
	return nil
}
