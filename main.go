package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type appConfig struct {
	WEBHOOK_URLS    string `envconfig:"WEBHOOK_URLS"`
	GITHUB_TOKEN    string `envconfig:"GITHUB_TOKEN"`
	REPO_TARGET     string `envconfig:"REPO_TARGET"`
	NUMBER_OF_ISSUE string `envconfig:"NUMBER_OF_ISSUE"`
	REPO_SOURCE     string `envconfig:"REPO_SOURCE"`
}

func main() {
	var appConfig appConfig
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}
	if err := envconfig.Process("", &appConfig); err != nil {
		fmt.Printf("Error processing environment variables: %v\n", err)
		return
	}

	webhooks := strings.Split(appConfig.WEBHOOK_URLS, ",")
	issueUrl := fmt.Sprintf("https://api.github.com/repos/%s/issues/%s", appConfig.REPO_TARGET, appConfig.NUMBER_OF_ISSUE)

	response, err := Fetch(issueUrl)
	if err != nil {
		fmt.Printf("Error fetching issue: %v\n", err)
		return
	}

	var body struct {
		Body string `json:"body"`
	}
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return
	}

	bodyInt, err := strconv.Atoi(body.Body)
	if err != nil {
		fmt.Printf("Error converting body to integer: %v\n", err)
		return
	}
	oldCheck := time.UnixMilli(int64(bodyInt))

	pullRequestsResponse, err := Fetch(fmt.Sprintf("https://api.github.com/repos/%s/pulls?state=all", appConfig.REPO_SOURCE))
	if err != nil {
		fmt.Printf("Error fetching pull requests: %v\n", err)
		return
	}

	var pullRequests PullRequests
	if err := json.NewDecoder(pullRequestsResponse.Body).Decode(&pullRequests); err != nil {
		fmt.Printf("Error decoding pull requests response: %v\n", err)
		return
	}

	slices.SortFunc(pullRequests, func(a, b PullRequest) int {
		return cmp.Compare(a.Number, b.Number)
	})

	for _, item := range pullRequests {
		createdAt := item.CreatedAt
		mergedAt := item.MergedAt
		closedAt := item.ClosedAt

		if createdAt.Before(oldCheck) && (mergedAt == nil || mergedAt.Before(oldCheck)) && (closedAt == nil || closedAt.Before(oldCheck)) {
			continue
		}

		embed := map[string]any{
			"author": map[string]any{
				"name":     item.User.Login,
				"url":      item.User.HTMLURL,
				"icon_url": item.User.AvatarURL,
			},
			"url": item.HTMLURL,
		}
		if item.Body != nil {
			embed["description"] = *item.Body
		}

		if createdAt.After(oldCheck) {
			embed["color"] = 0x4adb40
			embed["title"] = fmt.Sprintf("Pull request opened: #%d %s", item.Number, item.Title)

			if err := SendEmbed(embed, "created", webhooks); err != nil {
				fmt.Printf("Error sending embed for created pull request: %v\n", err)
			}
		}

		if mergedAt == nil && closedAt != nil && closedAt.After(oldCheck) {
			embed["color"] = 0xeb4034
			embed["title"] = fmt.Sprintf("Pull request closed: #%d %s", item.Number, item.Title)

			if err := SendEmbed(embed, "closed", webhooks); err != nil {
				fmt.Printf("Error sending embed for closed pull request: %v\n", err)
			}
		}

		if mergedAt != nil && mergedAt.After(oldCheck) {
			embed["color"] = 0x983ac7
			embed["title"] = fmt.Sprintf("Pull request merged: #%d %s", item.Number, item.Title)

			if err := SendEmbed(embed, "merged", webhooks); err != nil {
				fmt.Printf("Error sending embed for merged pull request: %v\n", err)
			}
		}

	}

	res, err := FetchWithBody("PATCH", issueUrl, map[string]any{
		"body": strconv.FormatInt(time.Now().UnixMilli(), 10),
	}, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", appConfig.GITHUB_TOKEN),
	})
	if err != nil {
		fmt.Printf("Error updating issue: %v\n", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code when updating issue: %d\n", res.StatusCode)
		return
	}
}
