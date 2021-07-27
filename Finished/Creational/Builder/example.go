package main

import "fmt"

func main() {
	var bldr = newNotificationBuilder()
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetSubTitle("This is a subtitle")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(5)
	bldr.SetMessage("This is a basic notification")
	bldr.SetType("alert")

	notif, _ := bldr.Build()
	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Printf("Notification: %+v\n", notif)
	}
}
