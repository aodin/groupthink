package reddit

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Reddit struct{}

func (reddit *Reddit) Search(u *url.URL) {}

func unmarshalSearch(content []byte) ([]Result, error) {
	results := struct {
		Data struct {
			Children []Result `json:"children"`
		} `json:"data"`
	}{}

	if err := json.Unmarshal(content, &results); err != nil {
		return nil, fmt.Errorf("Could not parse Reddit search: %s", err)
	}

	return results.Data.Children, nil
}

type Result struct {
	Kind string `json:"kind"`
	Data struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Score     int64     `json:"score"`
		Domain    string    `json:"domain"`
		Subreddit string    `json:"subreddit"`
		Created   Timestamp `json:"created_utc"`
	}
}
