package xkcdcom

import (
	"encoding/json"

	"net/http"
)

func Get(url string) (*Xkcd, error) {
	req, _ := http.NewRequest("GET", url, nil)
	//	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	var result Xkcd
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
