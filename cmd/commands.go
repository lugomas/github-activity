package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type GHEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
}

func RequestAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch data: %v\n", err)
	}

	defer resp.Body.Close()

	// Check if the user exists or there was another issue with the request
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: received status code %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetEvent(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	body, err := RequestAPI(url)
	if err != nil {
		log.Fatal(err)
	}

	var events []GHEvent
	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Fatal(err)
	}

	// Check if there are any events
	if len(events) == 0 {
		fmt.Printf("No recent activity found for user: %s\n", username)
		return
	}

	for _, event := range events {
		action := getActionFromType(event.Type)
		fmt.Printf("%s: %s at %s\n", action, event.Repo.Name, event.CreatedAt.Format(time.DateTime))
	}
}

// getActionFromType returns a human-readable action for each GitHub event type
func getActionFromType(eventType string) string {
	switch eventType {
	case "PushEvent":
		return "- Pushed to"
	case "IssuesEvent":
		return "- Opened an issue in"
	case "WatchEvent":
		return "- Starred"
	case "ForkEvent":
		return "- Forked"
	case "PullRequestEvent":
		return "- Created a pull request in"
	case "CreateEvent":
		return "- Created repository"
	default:
		return "- Unknown Event"
	}
}
