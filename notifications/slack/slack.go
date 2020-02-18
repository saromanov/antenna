package slack

import (
	"github.com/nlopes/slack"
	"github.com/saromanov/antenna/notifications"
	structs "github.com/saromanov/antenna/structs/v1"
)
type Slack struct {
	api *slack.Api
}

// New provides initialization of the Slack client
func New(token string) notifications.Notification {
	return &Slack {
		api: slack.Api(token),
	}
}

// Send provides sending of the message
func (s*Slack) Send(n *structs.Notification) error {
	return nil
}