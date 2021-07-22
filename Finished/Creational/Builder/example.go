package main

import "fmt"

func main() {
	var bldr = newNotificationBuilder()
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon 1")
	bldr.SetMessage("This is a basic notification")

	notif, _ := bldr.Build()
	fmt.Printf("Notification: %+v\n", notif)
}
