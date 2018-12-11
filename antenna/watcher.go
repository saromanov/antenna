package antenna

import (
	"fmt"
	"github.com/saromanov/antenna/container/docker"
)
// watcher creates object for watch running containers
// and gettign info from this
type watcher struct {
	dockerClient *docker.Docker
}

func (w*watcher) Watch() {
	containers, err := w.dockerClient.GetContainers(nil)
	if err != nil {
		fmt.Printf("unable to get list of containers: %v\n", err)
	}

	fmt.Println("Containers: ", containers)
}