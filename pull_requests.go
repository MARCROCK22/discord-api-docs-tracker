// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pullRequests, err := UnmarshalPullRequests(bytes)
//    bytes, err = pullRequests.Marshal()

package main

import (
	"encoding/json"
	"time"
)

type PullRequests []PullRequest

func UnmarshalPullRequests(data []byte) (PullRequests, error) {
	var r PullRequests
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PullRequests) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PullRequest struct {
	URL                string            `json:"url"`
	ID                 int64             `json:"id"`
	NodeID             string            `json:"node_id"`
	HTMLURL            string            `json:"html_url"`
	DiffURL            string            `json:"diff_url"`
	PatchURL           string            `json:"patch_url"`
	IssueURL           string            `json:"issue_url"`
	Number             int64             `json:"number"`
	State              State             `json:"state"`
	Locked             bool              `json:"locked"`
	Title              string            `json:"title"`
	User               UserClass         `json:"user"`
	Body               *string           `json:"body"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
	ClosedAt           *time.Time        `json:"closed_at"`
	MergedAt           *time.Time        `json:"merged_at"`
	MergeCommitSHA     string            `json:"merge_commit_sha"`
	Assignees          []UserClass       `json:"assignees"`
	RequestedReviewers []UserClass       `json:"requested_reviewers"`
	RequestedTeams     []any             `json:"requested_teams"`
	Labels             []Label           `json:"labels"`
	Milestone          any               `json:"milestone"`
	Draft              bool              `json:"draft"`
	CommitsURL         string            `json:"commits_url"`
	ReviewCommentsURL  string            `json:"review_comments_url"`
	ReviewCommentURL   string            `json:"review_comment_url"`
	CommentsURL        string            `json:"comments_url"`
	StatusesURL        string            `json:"statuses_url"`
	Head               Base              `json:"head"`
	Base               Base              `json:"base"`
	Links              Links             `json:"_links"`
	AuthorAssociation  AuthorAssociation `json:"author_association"`
	AutoMerge          any               `json:"auto_merge"`
	Assignee           *UserClass        `json:"assignee"`
	ActiveLockReason   any               `json:"active_lock_reason"`
}

type UserClass struct {
	Login             string       `json:"login"`
	ID                int64        `json:"id"`
	NodeID            string       `json:"node_id"`
	AvatarURL         string       `json:"avatar_url"`
	GravatarID        string       `json:"gravatar_id"`
	URL               string       `json:"url"`
	HTMLURL           string       `json:"html_url"`
	FollowersURL      string       `json:"followers_url"`
	FollowingURL      string       `json:"following_url"`
	GistsURL          string       `json:"gists_url"`
	StarredURL        string       `json:"starred_url"`
	SubscriptionsURL  string       `json:"subscriptions_url"`
	OrganizationsURL  string       `json:"organizations_url"`
	ReposURL          string       `json:"repos_url"`
	EventsURL         string       `json:"events_url"`
	ReceivedEventsURL string       `json:"received_events_url"`
	Type              Type         `json:"type"`
	UserViewType      UserViewType `json:"user_view_type"`
	SiteAdmin         bool         `json:"site_admin"`
}

type Base struct {
	Label string    `json:"label"`
	Ref   string    `json:"ref"`
	SHA   string    `json:"sha"`
	User  UserClass `json:"user"`
	Repo  Repo      `json:"repo"`
}

type Repo struct {
	ID                        int64                     `json:"id"`
	NodeID                    RepoNodeID                `json:"node_id"`
	Name                      RepoName                  `json:"name"`
	FullName                  FullName                  `json:"full_name"`
	Private                   bool                      `json:"private"`
	Owner                     UserClass                 `json:"owner"`
	HTMLURL                   string                    `json:"html_url"`
	Description               Description               `json:"description"`
	Fork                      bool                      `json:"fork"`
	URL                       string                    `json:"url"`
	ForksURL                  string                    `json:"forks_url"`
	KeysURL                   string                    `json:"keys_url"`
	CollaboratorsURL          string                    `json:"collaborators_url"`
	TeamsURL                  string                    `json:"teams_url"`
	HooksURL                  string                    `json:"hooks_url"`
	IssueEventsURL            string                    `json:"issue_events_url"`
	EventsURL                 string                    `json:"events_url"`
	AssigneesURL              string                    `json:"assignees_url"`
	BranchesURL               string                    `json:"branches_url"`
	TagsURL                   string                    `json:"tags_url"`
	BlobsURL                  string                    `json:"blobs_url"`
	GitTagsURL                string                    `json:"git_tags_url"`
	GitRefsURL                string                    `json:"git_refs_url"`
	TreesURL                  string                    `json:"trees_url"`
	StatusesURL               string                    `json:"statuses_url"`
	LanguagesURL              string                    `json:"languages_url"`
	StargazersURL             string                    `json:"stargazers_url"`
	ContributorsURL           string                    `json:"contributors_url"`
	SubscribersURL            string                    `json:"subscribers_url"`
	SubscriptionURL           string                    `json:"subscription_url"`
	CommitsURL                string                    `json:"commits_url"`
	GitCommitsURL             string                    `json:"git_commits_url"`
	CommentsURL               string                    `json:"comments_url"`
	IssueCommentURL           string                    `json:"issue_comment_url"`
	ContentsURL               string                    `json:"contents_url"`
	CompareURL                string                    `json:"compare_url"`
	MergesURL                 string                    `json:"merges_url"`
	ArchiveURL                string                    `json:"archive_url"`
	DownloadsURL              string                    `json:"downloads_url"`
	IssuesURL                 string                    `json:"issues_url"`
	PullsURL                  string                    `json:"pulls_url"`
	MilestonesURL             string                    `json:"milestones_url"`
	NotificationsURL          string                    `json:"notifications_url"`
	LabelsURL                 string                    `json:"labels_url"`
	ReleasesURL               string                    `json:"releases_url"`
	DeploymentsURL            string                    `json:"deployments_url"`
	CreatedAt                 time.Time                 `json:"created_at"`
	UpdatedAt                 time.Time                 `json:"updated_at"`
	PushedAt                  time.Time                 `json:"pushed_at"`
	GitURL                    GitURL                    `json:"git_url"`
	SSHURL                    SSHURL                    `json:"ssh_url"`
	CloneURL                  string                    `json:"clone_url"`
	SvnURL                    string                    `json:"svn_url"`
	Homepage                  string                    `json:"homepage"`
	Size                      int64                     `json:"size"`
	StargazersCount           int64                     `json:"stargazers_count"`
	WatchersCount             int64                     `json:"watchers_count"`
	Language                  *Language                 `json:"language"`
	HasIssues                 bool                      `json:"has_issues"`
	HasProjects               bool                      `json:"has_projects"`
	HasDownloads              bool                      `json:"has_downloads"`
	HasWiki                   bool                      `json:"has_wiki"`
	HasPages                  bool                      `json:"has_pages"`
	HasDiscussions            bool                      `json:"has_discussions"`
	ForksCount                int64                     `json:"forks_count"`
	MirrorURL                 any                       `json:"mirror_url"`
	Archived                  bool                      `json:"archived"`
	Disabled                  bool                      `json:"disabled"`
	OpenIssuesCount           int64                     `json:"open_issues_count"`
	License                   License                   `json:"license"`
	AllowForking              bool                      `json:"allow_forking"`
	IsTemplate                bool                      `json:"is_template"`
	WebCommitSignoffRequired  bool                      `json:"web_commit_signoff_required"`
	HasPullRequests           bool                      `json:"has_pull_requests"`
	PullRequestCreationPolicy PullRequestCreationPolicy `json:"pull_request_creation_policy"`
	Topics                    []Topic                   `json:"topics"`
	Visibility                UserViewType              `json:"visibility"`
	Forks                     int64                     `json:"forks"`
	OpenIssues                int64                     `json:"open_issues"`
	Watchers                  int64                     `json:"watchers"`
	DefaultBranch             DefaultBranch             `json:"default_branch"`
}

type License struct {
	Key    Key           `json:"key"`
	Name   LicenseName   `json:"name"`
	SpdxID SpdxID        `json:"spdx_id"`
	URL    string        `json:"url"`
	NodeID LicenseNodeID `json:"node_id"`
}

type Label struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}

type Links struct {
	Self           Comments `json:"self"`
	HTML           Comments `json:"html"`
	Issue          Comments `json:"issue"`
	Comments       Comments `json:"comments"`
	ReviewComments Comments `json:"review_comments"`
	ReviewComment  Comments `json:"review_comment"`
	Commits        Comments `json:"commits"`
	Statuses       Comments `json:"statuses"`
}

type Comments struct {
	Href string `json:"href"`
}

type Type string

type UserViewType string

type AuthorAssociation string

type DefaultBranch string

type Description string

type FullName string

type GitURL string

type Language string

type Key string

type LicenseName string

type LicenseNodeID string

type SpdxID string

type RepoName string

type RepoNodeID string

type PullRequestCreationPolicy string

type SSHURL string

type Topic string

type State string
