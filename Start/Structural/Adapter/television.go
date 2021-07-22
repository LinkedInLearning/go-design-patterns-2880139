package main

// This is the television interface we want to use with both TV types
type television interface {
	volumeUp() int
	volumeDown() int
	channelUp() int
	channelDown() int
	turnOn()
	turnOff()
	goToChannel(ch int)
}
