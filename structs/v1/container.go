package v1

import "time"

// Container define struct for container representation
type Container struct {
	CreatedTime time.Time `json:"created_time,omitempty"`
	Name        string    `json:"name"`
}
