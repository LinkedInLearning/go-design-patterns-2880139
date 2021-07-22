package main

import "fmt"

// Define the interface for an observer type
type observer interface {
	onUpdate(data string)
}

// Our DataListener observer will have a name
type DataListener struct {
	Name string
}

// To conform to the interface, it must have an onUpdate function
func (dl *DataListener) onUpdate(data string) {
	fmt.Println("Listener:", dl.Name, "got data change:", data)
}
