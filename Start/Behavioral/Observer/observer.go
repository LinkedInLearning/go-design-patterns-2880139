package main

// Define the interface for an observer type
type observer interface {
	onUpdate(data string)
}

// Our DataListener observer will have a name
type DataListener struct {
	Name string
}

// TODO: To conform to the interface, it must have an onUpdate function
