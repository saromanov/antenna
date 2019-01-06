package v1

// Notification defines payload for sending notifications
type Notification struct {
	ID            string
	Body          string
	Event         string
	ContainerName string
}
