package main
import "fmt"

// The NotificationBuilder has fields exported as well as a few methods
// to demonstrate
type NotificationBuilder struct {
	Title    string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(pri int) {
	nb.Priority = pri
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.NotType = notType
}

func (nb *NotificationBuilder) Build() (Notification, error) {
	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("You need to specify a subtitle when using an icon")
	}
	if nb.Priority > 5 {
		return nil, fmt.Errorf("Priority must be 0 to 5")
	}

	return Notification{
		title:    nb.Title,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
