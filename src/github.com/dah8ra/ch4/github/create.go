package github

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	//	"net/url"
)

func Login() (*http.Response, error) {
	//		req, _ := http.NewRequest("GET", SimpleUrl, nil)
	req, _ := http.NewRequest("GET", BasicAuthUrl, nil)
	//	req.SetBasicAuth("dah8ra", "nightmare@8746")
	client := new(http.Client)
	return client.Do(req)
}

const token = "0ae5d7f27f8da1a981e71424d57d6055cccc6767"

// SearchIssues queries the GitHub issue tracker.
func CreateIssues(jsonStr []byte) (*http.Response, error) { //(*IssuesSearchResult, error) {
	req, _ := http.NewRequest("POST", Url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "dah8ra:0ae5d7f27f8da1a981e71424d57d6055cccc6767")
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	return client.Do(req)
	//	q := url.QueryEscape(strings.Join(query, " "))
	//	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonStr))
	//	if err != nil {
	//		return nil, err
	//	}
	//	if resp.StatusCode != http.StatusOK {
	//		resp.Body.Close()
	//		return nil, fmt.Errorf("create query failed: %s", resp.Status)
	//	}

	//	var result IssuesSearchResult
	//	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	//		resp.Body.Close()
	//		return nil, err
	//	}
	//	resp.Body.Close()
	//	return &result, nil
}

func CreateIssuesWithAuthToken(jsonStr []byte) (*http.Response, error) {
	values := url.Values{}
	values.Add("access_token", token)

	//	req, _ := http.NewRequest("GET", BasicAuthUrl, bytes.NewBuffer(jsonStr))
	req, _ := http.NewRequest("GET", BasicAuthUrl, strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")

	//	jar, _ := cookiejar.New(nil)
	//	client := http.Client{Jar: jar}
	client := new(http.Client)

	return client.Do(req)
}
