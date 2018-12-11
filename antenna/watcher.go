package antenna

import (
	"fmt"

	"github.com/saromanov/antenna/container/docker"
)

// containerWatcher creates object for watch running containers
// and gettign info from this
type containerWatcher struct {
	dockerClient *docker.Docker
}

func (w *containerWatcher) Watch() {
	containers, err := w.dockerClient.GetContainers(nil)
	if err != nil {
		fmt.Printf("unable to get list of containers: %v\n", err)
	}

	fmt.Println("Containers: ", containers)
}
