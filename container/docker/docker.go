package docker

import (
	"github.com/fsouza/go-dockerclient"
)

const endpoint = "unix:///var/run/docker.sock"

// Docker provides implementation of the Docker
// logic
type Docker struct {
	client *Client
}

func Init() *Docker {
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	return Docker{client: client}
}
