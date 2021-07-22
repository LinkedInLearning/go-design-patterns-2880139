package main

type television interface {
	volumeUp() int
	volumeDown() int
	channelUp() int
	channelDown() int
	turnOn()
	turnOff()
	goToChannel(ch int)
}
