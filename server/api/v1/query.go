package v1

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aodin/groupthink/apis/reddit"
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
	args := strings.TrimSpace(q.Text)
	if args == "" {
		return nil, fmt.Errorf("Please provide a URL as the only argument")
	}

	if !strings.HasPrefix(strings.ToLower(args), "http") {
		args = "http://" + args
	}
	return url.Parse(args)
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

// TODO Convert to a non-global if state ever gets attached
var redditAPI = reddit.API{}

// Query converts Slack POST data to a response
func Query(w http.ResponseWriter, r *http.Request) {
	// Return all responses as plain text
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Parse the POST
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Could not parse request: %s", err), 400)
		return
	}

	q := fromRequest(r)
	defer logger(q, time.Now())

	// Parse the text as a url
	u, err := q.URL()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// TODO query mutliple APIs
	comments, err := redditAPI.Search(u)
	if err != nil {
		// TODO Distinguish between a 400 and a 500?
		http.Error(w, err.Error(), 400)
		return
	}

	if len(comments) < 1 {
		http.Error(w, "No groupthink was found for this link", 404)
		return
	}

	timeAgo := time.Now().UTC().Sub(comments[0].Created.AsTime()).Hours()
	var age string
	switch {
	case timeAgo < 1:
		age = "less than an hour ago"
	case timeAgo < 24:
		age = "less than a day ago"
	case timeAgo < (24 * 7):
		age = "less than a week ago"
	default:
		age = fmt.Sprint("%d days ago", timeAgo/24.0)
	}

	message := fmt.Sprintf(`>>> %s
%s (%s)`,
		comments[0].Body,
		comments[0].Permalink,
		age,
	)
	w.Write([]byte(message))
}
