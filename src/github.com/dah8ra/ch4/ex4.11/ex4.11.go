package main

import (
	"encoding/json"
	"fmt"

	"github.com/dah8ra/ch4/github"
)

var ticket = []github.Ticket{
	{
		Title:     "Found a first bug",
		Body:      "I'm having a problem with this.",
		Assignee:  "dah8ra",
		Milestone: 1,
		Labels:    "Label1",
	},
}

/*
{
		Title:     "Found a second bug",
		Body:      "I'm having a problem with this.",
		Assignee:  "dah8ra",
		Milestone: 1,
		Labels:    "Label1",
	},
*/
type Output struct {
	Message string `json:"message"`
}

func main() {
	input, err := json.MarshalIndent(ticket, "", "	")
	//	resp, err := github.CreateIssuesWithAuthToken(input)
	result, err := github.GetSingleIssue(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s\n", result.Id)
	//	for _, item := range result.Item {
	//		fmt.Printf("#%-5d %9.9s %.55s\n",
	//			item.Number, item.User.Login, item.Title)
	//	}
}
