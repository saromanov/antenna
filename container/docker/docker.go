package docker

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
	structs "github.com/saromanov/antenna/structs/v1"
)

const endpoint = "unix:///var/run/docker.sock"

// Docker provides implementation of the Docker
// logic
type Docker struct {
	client *Client
}

// Init provides initialization of the docker
func Init() *Docker {
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	d := &Docker{client: client}
	return d
}

// GetContainers returns list of containers
func (d *Docker) GetContainers() ([]*structs.Container, error) {
	containers, err := d.client.ListContainers()
	if err != nil {
		return nil, fmt.Errorf("unable to get list of containers: %v", err)
	}
}
