package main

func main() {
	// Construct two DataListener observers and
	// give each one a name
	listener1 := DataListener{
		Name: "Listener 1",
	}
	listener2 := DataListener{
		Name: "Listener 2",
	}

	// Create the DataSubject that the listeners will observe
	subj := &DataSubject{}
	// TODO: Register each listener with the DataSubject

	// TODO: Change the data in the DataSubject - this will cause the
	// onUpdate method of each listener to be called

	// TODO: Try to unregister one of the observers

}
