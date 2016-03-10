package github

import "time"


const Url = "https://api.github.com/repos/dah8ra/golangtraining/issues"
const IssueUrl = "https://api.github.com/repos/dah8ra/golangtraining/issues/"
const BaseUrl = "https://api.github.com/"


type Missing struct {
	Message string
	Errors  *Errors
}

type Errors struct {
	Resource string
	Field    string `json:"title"`
	Code     string `json:"missing_field"`
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string `json:"state"`
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
