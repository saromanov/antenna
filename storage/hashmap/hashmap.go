package hashmap

import (
	"github.com/saromanov/antenna/storage"
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
	return &hashmap{
		data: data,
	}, nil
}

// Add provides adding of stat
func (h *hashmap) Add(metrics *structs.ContainerStat) error {
	return nil
}

func (i *hashmap) Search(req *structs.ContainerStatSearch) ([]*structs.ContainerStat, error) {
	return nil, nil
}

func (i *hashmap) Aggregate(req *structs.AggregateSearchRequest) (*structs.AggregateSearchResponse, error) {
	return nil, nil
}
func (i *hashmap) Close() error {
	return nil
}
