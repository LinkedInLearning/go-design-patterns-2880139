package main

import "fmt"

func main() {
	// Create instances of the two TV types with some default values
	tv1 := &SammysangTV{
		currentChan:   13,
		currentVolume: 35,
		tvOn:          true,
	}
	tv2 := &SohneeTV{
		vol:     20,
		channel: 9,
		isOn:    true,
	}

	// Because the SohneeTV implements the "television" interface, we don't need an adapter
	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(68)
	tv2.turnOff()

	fmt.Println("--------------------")

	// We need to create a SammysangTV adapter for the SammysangTV class, however
	// because it has an interface that's different from the one we want to use
	ssAdapt := &sammysangAdapter{
		sstv: tv1,
	}
	ssAdapt.turnOn()
	ssAdapt.volumeUp()
	ssAdapt.volumeDown()
	ssAdapt.channelUp()
	ssAdapt.channelDown()
	ssAdapt.goToChannel(68)
	ssAdapt.turnOff()
}
