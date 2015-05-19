package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type API struct{}

type Comments []Comment

// TODO Sort?

func (reddit *API) Search(u *url.URL) (Comments, error) {
	resp, err := http.Get(searchURL(u))
	if err != nil {
		return nil, fmt.Errorf("Reddit search failed: %s", err)
	}
	defer resp.Body.Close()

	articles, err := unmarshalSearch(json.NewDecoder(resp.Body))
	if err != nil {
		return nil, err
	}

	if len(articles) == 0 {
		return nil, fmt.Errorf("This article has not been posted")
	}

	var comments Comments

	// TODO Query all the articles that have comments
	for _, article := range articles {
		if article.Data.NumberOfComments > 0 {
			commentResp, err := http.Get(commentURL(article.Data.ID))
			if err != nil {
				return nil, fmt.Errorf("Reddit comment query failed: %s", err)
			}
			defer commentResp.Body.Close()

			raw, err := unmarshalComments(json.NewDecoder(commentResp.Body))
			if err != nil {
				return nil, err
			}

			// Convert raw comments to complete comments by adding article
			// information
			for _, comment := range raw {
				comment.Data.Title = article.Data.Title
				comment.Data.Permalink = article.Data.Permalink
				comments = append(comments, comment.Data)
			}
		}
	}
	return nil, fmt.Errorf("There are no posts with comments available")
}

func commentURL(id string) string {
	return fmt.Sprintf(
		"http://www.reddit.com/comments/%s.json?sort=top&limit=1&depth=0",
		id,
	)
}

type RawComment struct {
	Kind string  `json:"kind"`
	Data Comment `json:"data"`
}

type Comment struct {
	ID      string    `json:"id"`
	Author  string    `json:"author"`
	Score   int64     `json:"score"`
	Body    string    `json:"body"`
	Created Timestamp `json:"created_utc"`

	// The following fields will be added from the article
	Title     string `json:"-"`
	Permalink string `json:"-"`
}

func unmarshalComments(decoder *json.Decoder) ([]RawComment, error) {
	// The Reddit API returns the original article and the comments as items
	// zero and one in an array, use raw messages
	var response []json.RawMessage
	if err := decoder.Decode(&response); err != nil {
		return nil, fmt.Errorf("Could not parse Reddit comments: %s", err)
	}

	if len(response) == 0 {
		return nil, fmt.Errorf("Article missing from Reddit comments")
	}

	// Don't bother to unmarshal the parent, we have all that information from
	// the original search query

	if len(response) < 2 {
		// TODO no comments isn't an error yet
		return nil, nil
	}

	var comments = struct {
		Kind string `json:"kind"`
		Data struct {
			Children []RawComment `json:"children"`
		} `json:"data"`
	}{}

	// TODO or put parent and comment into a interface slice?
	if err := json.Unmarshal(response[1], &comments); err != nil {
		return nil, fmt.Errorf("Could not parse Reddit comment: %s", err)
	}
	return comments.Data.Children, nil
}

func searchURL(dirty *url.URL) string {
	// Search only the domain and path, drop scheme, query, and fragment
	return fmt.Sprintf(
		"http://www.reddit.com/search/json?q=url:%s",
		(&url.URL{Host: dirty.Host, Path: dirty.Path}).String(),
	)
}

type Article struct {
	Kind string `json:"kind"`
	Data struct {
		ID               string    `json:"id"`
		Title            string    `json:"title"`
		Score            int64     `json:"score"`
		Domain           string    `json:"domain"`
		Subreddit        string    `json:"subreddit"`
		Permalink        string    `json:"permalink"`
		Created          Timestamp `json:"created_utc"`
		NumberOfComments int64     `json:"num_comments"`
	} `json:"data"`
}

func unmarshalSearch(decoder *json.Decoder) ([]Article, error) {
	results := struct {
		Data struct {
			Children []Article `json:"children"`
		} `json:"data"`
	}{}

	if err := decoder.Decode(&results); err != nil {
		return nil, fmt.Errorf("Could not parse Reddit search: %s", err)
	}

	return results.Data.Children, nil
}
