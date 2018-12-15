package antenna

import (
	"fmt"

	"github.com/robfig/cron"
	"github.com/saromanov/antenna/container/docker"
)

// containerWatcher creates object for watch running containers
// and gettign info from this
type containerWatcher struct {
	dockerClient *docker.Docker
}

func (w *containerWatcher) Watch() {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		w.getContainers()
	})
}

func (w *containerWatcher) getContainers() {
	containers, err := w.dockerClient.GetContainers(nil)
	if err != nil {
		fmt.Printf("unable to get list of containers: %v\n", err)
		return
	}
	fmt.Println("Containers: ", containers)
}
