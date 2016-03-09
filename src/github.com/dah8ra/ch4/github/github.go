package github

import "time"

//const Url = "https://api.github.com/repos/dah8ra/golangtraining/issues"
const Url = "https://api.github.com/repos/dah8ra/golangtraining/issues"
const ClientId = "<your client key>"
const ClientSecret = "<your client secret>"

//const AuthorizationBaseUrl = "https://github.com/login/oauth/authorize"
const BasicAuthUrl = "https://api.github.com/user?0ae5d7f27f8da1a981e71424d57d6055cccc6767"
const SingleIssueUrl = "https://api.github.com/repos/dah8ra/golangtraining/issues/1"

//const OauthUrl = "https://api.github.com?access_token=OAUTH-TOKEN"
const BaseUrl = "https://api.github.com/"
const TokenUrl = "https://github.com/login/oauth/access_token"

type Ticket struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Assignee  string `json:"assignee"`
	Milestone int    `json:"milestone"`
	Labels    string `json:"labels"`
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type SingleIssueResult struct {
	Id        int
	Number    int
	Title     string
	Login     string
	State     string
	Assignee  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
