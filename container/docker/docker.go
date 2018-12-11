package docker

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
	structs "github.com/saromanov/antenna/structs/v1"
)

const defaultEndpoint = "unix:///var/run/docker.sock"

// Docker provides implementation of the Docker
// logic
type Docker struct {
	client *docker.Client
}

// Init provides initialization of the docker
func Init(conf *structs.ClientContainerConfig) *Docker {
	endpoint := conf.Endpoint
	if endpoint == "" {
		endpoint = defaultEndpoint
	}
	conf.Endpoint = endpoint
	client, err := createDockerClient(conf)
	if err != nil {
		panic(err)
	}
	d := &Docker{client: client}
	return d
}

// createDockerClient provides init of the docker client with TLS config or without
func createDockerClient(conf *structs.ClientContainerConfig) (*docker.Client, error) {
	if conf.CertPathEnv != "" {
		path := os.Getenv("DOCKER_CERT_PATH")
		ca := fmt.Sprintf("%s/ca.pem", path)
		cert := fmt.Sprintf("%s/cert.pem", path)
		key := fmt.Sprintf("%s/key.pem", path)
		return docker.NewTLSClient(conf.Endpoint, cert, key, ca)
	}
	return docker.NewClient(conf.Endpoint)
}

// GetContainers returns list of containers
func (d *Docker) GetContainers(opt *structs.ListContainersOptions) ([]*structs.Container, error) {
	containers, err := d.client.ListContainers(docker.ListContainersOptions{
		All:     opt.All,
		Size:    opt.Size,
		Limit:   opt.Limit,
		Before:  opt.Before,
		Since:   opt.Since,
		Filters: opt.Filters,
		Context: opt.Context,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get list of containers: %v", err)
	}
	cont, err := d.toContainerList(containers)
	if err != nil {
		return nil, err
	}
	return cont, nil

}

// Name returns name of container type
func (d *Docker) Name() string {
	return "docker"
}

// Start provides starting of container by id
func (d *Docker) Start(id string) error {
	return d.client.StartContainer(id, &docker.HostConfig{})
}

// Version returns current version of Docker API
func (d *Docker) Version() (string, error) {
	ver, err := d.client.Version()
	if err != nil {
		return "", fmt.Errorf("unable to get Docker version: %v", err)
	}
	return ver.Get("version"), nil
}

// toContainerList returns containers at inner representation
func (d *Docker) toContainerList(cl []docker.APIContainers) ([]*structs.Container, error) {
	containers := make([]*structs.Container, len(cl))
	for i, cont := range cl {
		containers[i] = d.toContainer(cont)
	}
	return containers, nil
}

// toContainer retrurns container at inner representation
func (d *Docker) toContainer(c docker.APIContainers) *structs.Container {
	return &structs.Container{
		Image:  c.Image,
		Names:  c.Names,
		Status: c.Status,
		State:  c.State,
	}
}
