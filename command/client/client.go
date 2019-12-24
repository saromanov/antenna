package client

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func Get(url string, response interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("unable to unmarshal response")
	}

	return nil
}