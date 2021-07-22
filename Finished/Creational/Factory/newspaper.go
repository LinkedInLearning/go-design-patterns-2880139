package main

import "fmt"

// Define a newspaper type and embed the publication interface
type newspaper struct {
	publication
}

// Define a Stringer interface that gives a string representation of the type
func (n newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

func createNewspaper(name string, pages int, publisher string) iPublication {
	return &newspaper{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
