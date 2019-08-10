package hashmap

import (
	structs "github.com/saromanov/antenna/structs/v1"
)

// hashmap represents local storage
// in the cases when remote storage in not available
type hashmap struct {
	data map[string]*structs.ContainerStat
}

// New creates storage based on name
// At the init stage
func New(conf *storage.Config) (storage.Storage, error) {
	return new(conf)
}

func new(conf *storage.Config) (storage.Storage, error) {
	data := map[string]*structs.ContainerStat{}
	return hashmap{
		data: data,
	}, nil
}