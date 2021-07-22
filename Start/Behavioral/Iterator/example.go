package main

import "fmt"

// Iterate using a callback function
func main() {
	// TODO: use IterateBooks to iterate via a callback function


	// TODO: Use IterateBooks to iterate via anonymous function


	// TODO: create a BookIterator
}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
