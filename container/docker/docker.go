package docker

// Docker provides implementation of the Docker
// logic
type Docker struct {
	client *Client
}

func Init() *Docker {
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	return Docker{client: client}
}
