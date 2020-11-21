package repository

import "time"

// CommitResponse is the response you get when getting a single commit from the
// Github API.
type CommitResponse struct {
	Author      CommitResponseAuthor    `json:"author"`
	Commit      Commit                  `json:"commit"`
	Committer   CommitResponseCommitter `json:"committer"`
	CommentsURL string                  `json:"comments_url"`
	Files       CommitFile              `json:"files"`
	HTMLURL     string                  `json:"html_url"`
	NodeID      string                  `json:"node_id"`
	Parents     []CommitParent          `json:"parents"`
	SHA         string                  `json:"sha"`
	Stats       CommitStats             `json:"stats"`
	URL         string                  `json:"url"`
}

// Commit is a single point in the Git history. It can be used when retrieving a
// single commit from a Github Repository.
type Commit struct {
	URL string `json:"url"`

	Message string `json:"message"`

	CommitTree         CommitTree         `json:"tree"`
	CommitAuthor       CommitAuthor       `json:"author"`
	CommitComitter     CommitComitter     `json:"comitter"`
	CommitVerification CommitVerification `json:"verification"`

	CommentCount int `json:"comment_count"`
}

// CommitAuthor is the author of the commit.
type CommitAuthor struct {
	Date  time.Time `json:"date"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

// CommitComitter is the commiter of the commit.
type CommitComitter struct {
	Date  time.Time `json:"date"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

// CommitTree contains a SHA hash and a URL.
type CommitTree struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

// CommitVerification describes the result of verifying the commit's signature.
//
// https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/about-commit-signature-verification
type CommitVerification struct {
	Payload   string `json:"payload"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Verified  bool   `json:"verified"`
}

// CommitResponseAuthor is the author of the commit.
type CommitResponseAuthor struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	NodeID    string `json:"node_id"`
	SiteAdmin bool   `json:"site_admin"`
	Type      string `json:"type"`

	AvatarURL         string `json:"avatar_url"`
	EventsURL         string `json:"events_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	GravatarID        string `json:"gravatar_id"`
	HTMLURL           string `json:"html_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	ReposURL          string `json:"repos_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	URL               string `json:"url"`
}

// CommitResponseCommitter is the commiter of the commit.
type CommitResponseCommitter struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	NodeID    string `json:"node_id"`
	SiteAdmin bool   `json:"site_admin"`
	Type      string `json:"type"`

	AvatarURL         string `json:"avatar_url"`
	EventsURL         string `json:"events_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	GravatarID        string `json:"gravatar_id"`
	HTMLURL           string `json:"html_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	ReposURL          string `json:"repos_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	URL               string `json:"url"`
}

// CommitParent contains information about the parent commit.
type CommitParent struct {
	URL string `json:"url"`
	SHA string `json:"sha"`
}

// CommitStats contains information about the size of the commit.
type CommitStats struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}

// CommitFile contains information about a file being committed.
type CommitFile struct {
	Additions int    `json:"additions"`
	BlobURL   string `json:"blob_url"`
	Changes   int    `json:"changes"`
	Deletions int    `json:"deletions"`
	Filename  string `json:"filename"`
	Patch     string `json:"patch"`
	Status    string `json:"status"`
	RawURL    string `json:"raw_url"`
}

// CommitRequest has the information for requesting a commit.
type CommitRequest struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
	Ref   string `json:"ref"`
}

// Path returns the URL Path for the request. It follows the following pattern:
//      /repos/{owner}/{repo}/commits/{ref}
func (c CommitRequest) Path() string {
	return "/repos/" + c.Owner + "/" + c.Repo + "/commits/" + c.Ref
}
