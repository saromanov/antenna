package antenna

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/saromanov/antenna/config"
	"github.com/saromanov/antenna/container/docker"
	"github.com/saromanov/antenna/storage"
	structs "github.com/saromanov/antenna/structs/v1"
	log "github.com/sirupsen/logrus"
)

// Application provides definition of the main
// interface for app
type Application struct {
	HTTPClient     http.Client
	Store          storage.Storage
	MapStore       storage.Storage
	Config         *config.Config
	events         chan *ContainerEvent
	dockerClient   *docker.Docker
	watcher        *containerWatcher
	containers     map[string]*structs.Container
	allContainers  map[string]*structs.Container
	containersLock *sync.RWMutex
	startTime      time.Time
	staticHostInfo *HostInfo
}

type containerInfo struct {
	Name string
	ID   string
}

// ContainerEventType provides definition for container event handling
type ContainerEventType int

const (
	// ContainerAdd defines event for adding a new container
	ContainerAdd ContainerEventType = iota + 1
	// ContainerRemove defines event for removing old one container
	ContainerRemove
	// ListContainers returns list of running containers
	ListContainers
)

// ContainerEvent event defines events on containers
type ContainerEvent struct {
	event      ContainerEventType
	containers []*structs.Container
	name       string
}

// Start provides starting of the app
func (a *Application) Start() error {
	var err error
	a.startTime = time.Now().UTC()
	a.events = make(chan *ContainerEvent)
	a.containersLock = &sync.RWMutex{}
	a.containers = make(map[string]*structs.Container)
	a.allContainers = make(map[string]*structs.Container)
	client := docker.Init(&structs.ClientContainerConfig{})
	a.dockerClient = client
	a.watcher = &containerWatcher{
		dockerClient: a.dockerClient,
		events:       a.events,
	}

	a.staticHostInfo, err = getStaticHostInfo()
	if err != nil {
		return fmt.Errorf("unable to get host info: %v", err)
	}
	go a.watcher.Watch()
	a.startEventWatcher()
	return nil
}

// getContainers returns map of containers
func (a *Application) getContainers() map[string]*structs.Container {
	return a.containers
}

// addContainer creating of the event after adding of the new container
func (a *Application) addContainer() {
	fmt.Println("Adding container")
}

func (a *Application) removeContainer(name string) {
	func() {
		a.containersLock.RLock()
		defer a.containersLock.RUnlock()
		delete(a.containers, name)
	}()
}

func (a *Application) processListContainers(containers []*structs.Container) {
	a.containersLock.RLock()
	defer a.containersLock.RUnlock()
	numOld := len(a.containers)
	old := copyMap(a.containers)
	for _, c := range containers {
		container, _ := a.dockerClient.Get(c.ID)
		stats := a.dockerClient.GetStats(container.ID)
		stats.Image = container.Image
		stats.Name = container.Name
		if err := a.insertStats(stats); err != nil {
			log.WithFields(log.Fields{"method": "processListContainers"}).Infof("unable to insert stat to the storage: %v", err)
		}
		var name string
		if len(c.Names) > 0 {
			name = c.Names[0]
		}
		a.containers[name] = c
		a.allContainers[name] = c
	}

	go func(p map[string]*structs.Container, cont []*structs.Container, numOldContainers int) {
		if numOldContainers > len(cont) {
			for _, c := range old {
				found := false
				for _, c2 := range cont {
					if c.ID == c2.ID {
						found = true
						break
					}
				}
				if !found {
					a.events <- &ContainerEvent{
						event: ContainerRemove,
						name:  c.Names[0],
					}
					log.WithFields(log.Fields{"method": "processListContainers"}).Infof("container was removed: %v", c.ID)
				}
			}
		}
		return
	}(old, containers, numOld)
}

// insertStat provides inserting of the container stat to the storage
// if storage is not available
func (a *Application) insertStats(stat *structs.ContainerStat) error {
	if a.Store == nil {
		return errors.New("storage is not defined")
	}
	err := a.Store.Add(stat)
	if err != nil {
		log.WithFields(log.Fields{"method": "processListContainers"}).Infof("unable to insert stat to the storage: %v", err)
		err = a.insertStatsToMap(stat)
		if err != nil {
			return fmt.Errorf("unable to insert stats to the temp storage: %v", err)
		}
		log.WithFields(log.Fields{"method": "processListContainers"}).Infof("stats was inserted to the temp storage")
	}

	oldData, err := a.getStatsFromMap()
	if err != nil {
		return err
	}
	for _, d := range oldData {
		if err := a.Store.Add(d); err != nil {
			log.WithFields(log.Fields{"method": "processListContainers"}).WithError(err).Errorf("new event was inserted")
			return err
		}
		log.WithFields(log.Fields{"method": "processListContainers"}).Infof("new event was inserted")
	}

	a.MapStore.Close()
	return nil
}

// insertStatsToMap probvides inserting to the temp map
// until main storage will be available
func (a *Application) insertStatsToMap(stat *structs.ContainerStat) error {
	if a.MapStore == nil {
		return errors.New("map storage is empty")
	}
	return a.MapStore.Add(stat)
}

// getStatsFromMap returns data from temp hashmap
func (a *Application) getStatsFromMap() ([]*structs.ContainerStat, error) {
	if a.MapStore == nil {
		return nil, errors.New("map storage is empty")
	}

	return a.MapStore.Search(&structs.ContainerStatSearch{})
}

func (a *Application) startEventWatcher() {
	for {
		select {
		case event := <-a.events:
			switch event.event {
			case ContainerAdd:
				a.addContainer()
			case ContainerRemove:
				a.removeContainer(event.name)
			case ListContainers:
				a.processListContainers(event.containers)
			}

		}
	}
}

type antenna struct {
	store      storage.Storage
	httpClient *http.Client
}

// New provides initialization on the app
func New() (*Application, error) {
	return &Application{
		Store:          nil,
		containersLock: &sync.RWMutex{},
		containers:     make(map[string]*structs.Container),
	}, nil
}

// GetContainerInfo returns info about running container
func (a *Application) GetContainerInfo(name string) (*structs.Container, error) {
	var cont *structs.Container
	var ok bool
	func() {
		a.containersLock.RLock()
		defer a.containersLock.RUnlock()
		cont, ok = a.containers[name]
	}()
	if !ok {
		return nil, fmt.Errorf("unknown container %q", name)
	}
	return cont, nil
}

// copyMap provides copy of the containers map into the new one
func copyMap(s map[string]*structs.Container) map[string]*structs.Container {
	response := make(map[string]*structs.Container)
	for k, v := range s {
		response[k] = v
	}
	return response
}
