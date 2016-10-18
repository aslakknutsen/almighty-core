package remoteworkitem

import (
	"encoding/json"

	"github.com/andygrunwald/go-jira"
)

// Jira represents the Jira tracker provider
type Jira struct {
	URL   string
	Query string
}

// Fetch collects data from Jira
func (j *Jira) Fetch() chan map[string]string {
	item := make(chan map[string]string)
	go func() {
		client, _ := jira.NewClient(nil, j.URL)
		issues, _, _ := client.Issue.Search(j.Query, nil)
		for l := range issues {
			i := make(map[string]string)
			id, _ := json.Marshal(issues[l].Key)
			issue, _, _ := client.Issue.Get(issues[l].Key)
			content, _ := json.Marshal(issue)
			i = map[string]string{"id": string(id), "content": string(content)}
			item <- i
		}
		close(item)
	}()
	return item
}
