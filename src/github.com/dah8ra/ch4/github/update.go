package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateIssues(url string, jsonStr []byte) (*Issue, error) {
	//	req, _ := http.NewRequest("PATCH", url, bytes.NewBufferString(jsonStr))
	req, _ := http.NewRequest("PATCH", url, bytes.NewReader(jsonStr))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var missing Missing
		if err := json.NewDecoder(resp.Body).Decode(&missing); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		return nil, fmt.Errorf("%s", missing.Message)
		//		return nil, fmt.Errorf("create query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
