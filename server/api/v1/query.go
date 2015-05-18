package v1

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type slackQuery struct {
	Token       string
	TeamID      string
	TeamDomain  string
	ChannelID   string
	ChannelName string
	UserID      string
	UserName    string
	Command     string
	Text        string
}

func (q slackQuery) URL() (*url.URL, error) {
	// TODO error if the command was wrong
	// TODO error if the text was not a url
	return nil, fmt.Errorf("A work in progress")
}

func logger(q slackQuery, then time.Time) {
	if q.Command == "" {
		// TODO separate failure logs?
		log.Println("failed query")
		return
	}
	log.Printf(
		"%s (%s: %s) on %s ran %s (%s) for %s",
		q.UserName,
		q.TeamDomain,
		q.Token,
		q.ChannelName,
		q.Command,
		q.Text,
		time.Now().Sub(then),
	)
}

func fromRequest(r *http.Request) (slack slackQuery) {
	slack.Token = r.FormValue("token")
	slack.TeamID = r.FormValue("team_id")
	slack.TeamDomain = r.FormValue("team_domain")
	slack.ChannelID = r.FormValue("channel_id")
	slack.ChannelName = r.FormValue("channel_name")
	slack.UserID = r.FormValue("user_id")
	slack.UserName = r.FormValue("user_name")
	slack.Command = r.FormValue("command")
	slack.Text = r.FormValue("text")
	return
}

// Query converts Slack POST data to a response
func Query(w http.ResponseWriter, r *http.Request) {
	// Parse the POST
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Could not parse request: %s", err), 400)
		return
	}

	q := fromRequest(r)
	defer logger(q, time.Now())

	// Parse the text as a url
	_, err := q.URL()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// TODO Query the requested APIs async and with timeouts
	http.Error(w, "In progress", 400)
}
