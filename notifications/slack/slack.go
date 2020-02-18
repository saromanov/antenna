package slack

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/saromanov/antenna/notifications"
	structs "github.com/saromanov/antenna/structs/v1"
)

type Slack struct {
	api       *slack.Client
	channelID string
}

// New provides initialization of the Slack client
func New(token, channelID string) notifications.Notifications {
	return &Slack{
		api:       slack.New(token),
		channelID: channelID,
	}
}

// Send provides sending of the message
func (s *Slack) Send(n *structs.Notification) error {
	channelID, timestamp, err := s.api.PostMessage("CHANNEL_ID", slack.MsgOptionText(n.Body, false), nil)
	if err != nil {
		return fmt.Errorf("unable to post message: %v", err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
