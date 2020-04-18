package github

import "time"

// IssuesURL link
const IssuesURL =  "https://api.github.com/search/issues"

// IssuesSearchResult structure
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue structure
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time
}

// User structure
type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}





