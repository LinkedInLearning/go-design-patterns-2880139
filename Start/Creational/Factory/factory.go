package main

import "fmt"

// newPublication is a factory function that creates the specified publication type
func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	// TODO: Create the right kind of publication based on the given type

	return nil, fmt.Errorf("No such publication type")
}
