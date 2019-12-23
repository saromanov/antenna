package antenna

import (
	"github.com/robfig/cron"
	"github.com/saromanov/antenna/container/docker"
	structs "github.com/saromanov/antenna/structs/v1"
	log "github.com/sirupsen/logrus"
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
	log.Infof("starting of the watcher")
	c.Start()
}

func (w *containerWatcher) getContainers() []*structs.Container {
	containers, err := w.dockerClient.List(nil)
	if err != nil {
		log.WithFields(log.Fields{
			"stage": "watcher",
		}).WithError(err).Errorf("unable to get list of containers")
		return nil
	}
	if len(containers) == 0 {
		log.Infof("unable to find containers")
	}
	return containers
}
