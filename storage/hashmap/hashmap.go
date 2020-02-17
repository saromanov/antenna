package hashmap

import (
	"github.com/saromanov/antenna/storage"
	structs "github.com/saromanov/antenna/structs/v1"
	uuid "github.com/satori/go.uuid"
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
	u := uuid.NewV4()
	h.data[u.String()] = metrics
	return nil
}

func (h *hashmap) Search(req *structs.ContainerStatSearch) ([]*structs.ContainerStat, error) {
	response := []*structs.ContainerStat{}
	for _, value := range h.data {
		response = append(response, value)
	}
	return response, nil
}

func (h *hashmap) Aggregate(req *structs.AggregateSearchRequest) (*structs.AggregateSearchResponse, error) {
	return nil, nil
}
func (h *hashmap) Close() error {
	h.data = map[string]*structs.ContainerStat{}
	return nil
}

func (i *hashmap) Info() map[string]interface{} {
	return map[string]interface{}{
		"name": "hashmap",
	}
}
