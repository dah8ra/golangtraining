package github

import (
	"encoding/json"
	"net/http"
)

func GetSingleIssue(jsonStr []byte) (*SingleIssueResult, error) {
	req, _ := http.NewRequest("GET", SingleIssueUrl, nil)
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, _ := client.Do(req)
	var result SingleIssueResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
