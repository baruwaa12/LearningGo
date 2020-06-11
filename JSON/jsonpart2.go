package main
package github

import (
	"time"
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const IssuesURL =  "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount string 'json:"total_count'
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string 'json:"html_url"'
	Title     string
	State     string
	User      *User
	CreatedAt time.Time 'json:"created_at"'
	Body      string
}

type User struct {
	Login    string
	HTMLURL  string 'json:"html_url"'
}

func main() {
	fmt.PrintLn("Running")
}



