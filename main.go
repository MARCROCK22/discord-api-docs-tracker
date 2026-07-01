package main

import (
	"cmp"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	webhookURLs := flag.String("webhooks", "", "comma-separated Discord webhook URLs")
	githubToken := flag.String("token", "", "GitHub token with permission to edit the state issue")
	repoTarget := flag.String("repo-target", "", "owner/repo holding the state issue")
	repoSource := flag.String("repo-source", "", "owner/repo to watch for pull requests")
	numberOfIssue := flag.String("issue", "", "issue number storing the last-check timestamp")
	flag.Parse()

	if *webhookURLs == "" || *githubToken == "" || *repoTarget == "" || *repoSource == "" || *numberOfIssue == "" {
		fmt.Println("missing required flags")
		flag.Usage()
		return
	}

	webhooks := strings.Split(*webhookURLs, ",")
	issueUrl := fmt.Sprintf("https://api.github.com/repos/%s/issues/%s", *repoTarget, *numberOfIssue)

	response, err := Fetch(issueUrl, *githubToken)
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

	pullRequestsResponse, err := Fetch(fmt.Sprintf("https://api.github.com/repos/%s/pulls?state=all", *repoSource), *githubToken)
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
		if item.User.Login == "dependabot[bot]" {
			continue
		}

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
			desc := *item.Body
			if r := []rune(desc); len(r) > 4096 {
				desc = string(r[:4093]) + "..."
			}
			embed["description"] = desc
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
		"Authorization": fmt.Sprintf("Bearer %s", *githubToken),
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
