package notifications

import structs "github.com/saromanov/antenna/structs/v1"

// Notifications defines object for sending notifications
// based on events
type Notifications interface {
	Send(*structs.Notification) error
}
