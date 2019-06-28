package antenna

import (
	"fmt"

	"github.com/robfig/cron"
	"github.com/saromanov/antenna/container/docker"
	structs "github.com/saromanov/antenna/structs/v1"
)

// containerWatcher creates object for watch running containers
// and getting info from this
type containerWatcher struct {
	dockerClient *docker.Docker
	events       chan *ContainerEvent
}

func (w *containerWatcher) Watch() {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		w.events <- &ContainerEvent{
			event:      ListContainers,
			containers: w.getContainers(),
		}
	})
	c.Start()
}

func (w *containerWatcher) getContainers() []*structs.Container {
	containers, err := w.dockerClient.List(nil)
	if err != nil {
		fmt.Printf("unable to get list of containers: %v\n", err)
		return nil
	}
	return containers
}
