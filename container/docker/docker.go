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
	if conf == nil {
		panic("Client container Config is nil")
	}
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

// List returns list of containers
func (d *Docker) List(opt *structs.ListContainersOptions) ([]*structs.Container, error) {
	dopt := docker.ListContainersOptions{}
	if opt != nil {
		dopt = docker.ListContainersOptions{
			All:     opt.All,
			Size:    opt.Size,
			Limit:   opt.Limit,
			Before:  opt.Before,
			Since:   opt.Since,
			Filters: opt.Filters,
			Context: opt.Context,
		}
	}
	containers, err := d.client.ListContainers(dopt)
	if err != nil {
		return nil, fmt.Errorf("unable to get list of containers: %v", err)
	}
	cont, err := d.toContainerList(containers)
	if err != nil {
		return nil, err
	}
	return cont, nil

}

// GetStats returns stat for container
func (d *Docker) GetStats(id string) *structs.ContainerStat {
	statsC := make(chan *docker.Stats)
	go func() {
		if err := d.client.Stats(docker.StatsOptions{ID: id, Stats: statsC, Stream: false}); err != nil {
			fmt.Println("unable to get stats: ", err)
		}
	}()
	stats := <-statsC
	resp := &structs.ContainerStat{
		CPU: structs.CPUStat{
			TotalUsage:     stats.CPUStats.CPUUsage.TotalUsage,
			OnlineCPUs:     stats.CPUStats.OnlineCPUs,
			SystemCPUUsage: stats.CPUStats.SystemCPUUsage,
		},
		NumProcs: stats.NumProcs,
		Cache:    stats.MemoryStats.Stats.Cache,
		MaxUsage: stats.MemoryStats.MaxUsage,
		Limit:    stats.MemoryStats.Limit,
		Usage:    stats.MemoryStats.Usage,
	}
	fmt.Println(stats.CPUStats.CPUUsage.PercpuUsage)
	return resp
}

// Name returns name of container type
func (d *Docker) Name() string {
	return "docker"
}

// Start provides starting of container by id
func (d *Docker) Start(opt *structs.StartContainerOptions) error {
	if opt.ID == "" {
		return errIDNotDefined
	}
	return d.client.StartContainer(opt.ID, &docker.HostConfig{})
}

// Get provides getting of container data by id
func (d *Docker) Get(id string) (*structs.Container, error) {
	if id == "" {
		return nil, errIDNotDefined
	}
	container, err := d.client.InspectContainer(id)
	if err != nil {
		return nil, fmt.Errorf("unable to inspect container: %v", err)
	}
	return d.fromInspectContainer(container), nil
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
		ID:         c.ID,
		Image:      c.Image,
		Names:      c.Names,
		Status:     c.Status,
		State:      c.State,
		SizeRw:     c.SizeRw,
		SizeRootFs: c.SizeRootFs,
		Labels:     c.Labels,
	}
}

// fromInspectContainer provides converting of Conteiner type from Docker
// library into inner representation of container
func (d *Docker) fromInspectContainer(c *docker.Container) *structs.Container {
	return &structs.Container{
		ID:           c.ID,
		Image:        c.Image,
		Name:         c.Name,
		Status:       c.State.Status,
		CreatedTime:  c.Created,
		RestartCount: c.RestartCount,
		Running:      c.State.Running,
		Paused:       c.State.Paused,
		Restarting:   c.State.Restarting,
		OOMKilled:    c.State.OOMKilled,
		Error:        c.State.Error,
		Args:         c.Args,
	}
}

func (d *Docker) toStats(s *docker.Stats) *structs.ContainerStat {
	return &structs.ContainerStat{}
}
