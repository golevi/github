/*
Package github provides structs and methods for interacting with the Github API.

Commits

Create a CommitRequest and start building a new request using the CommitRequest
path.

	cr := CommitRequest{
		Owner: "levidurfee",
		Repo:  "test",
		Ref:   "abcdefgh12345789ijklmnop",
	}

	req := http.NewRequest("GET", cr.Path(), nil)

https://docs.github.com/en/free-pro-team@latest/rest/reference/repos#get-a-commit
*/
package github
