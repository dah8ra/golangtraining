package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	"github.com/dah8ra/ch4/github"
)

var issue github.Issue

var n = flag.Bool("n", false, "omit trailing newline")
var read = flag.Bool("r", false, "read ticket")
var create = flag.Bool("c", false, "create new ticket")
var ticketTitle = flag.String("t", "default", "tiket title")
var update = flag.Bool("u", false, "update ticket")
var done = flag.Bool("d", false, "close ticket")
var num = flag.Int("num", 0, "ticket number")

func main() {
	flag.Parse()
	url := createIssueUrl(*num)
	fmt.Printf("-------> %s\n", url)

	if *read {
		result, _ := github.GetSingleIssue(url)
		fmt.Printf("#%-5d %9.9s %.55s\n", result.Number, result.User.Login, result.Title)
		return
	}

	issue.Title = *ticketTitle
	if *done {
		issue.State = "close"
	}

	input, _ := json.MarshalIndent(issue, "", "	")
	if *create {
		result, _ := github.CreateIssues(input)
		fmt.Printf("#%-5d %9.9s %.55s\n", result.Number, result.User.Login, result.Title)
	} else if *update {
		result, _ := github.UpdateIssues(url, input)
		fmt.Printf("#%-5d %9.9s %.55s\n", result.Number, result.User.Login, result.Title)
	}

}

/*
func main() {
	flag.Parse()
	if *read {
		url := createIssueUrl(*num)
		result, _ := github.GetSingleIssue(url)
		fmt.Printf("%s\n", result.Title)
	} else if *create {
		input := createJson(*ticketTitle)
		github.CreateIssues(input)
	} else if *update {
		url := createIssueUrl(*num)
		temp := createJsonFormat(*ticketTitle)
		input := createJson(temp)
		result, _ := github.UpdateIssues(url, *ticketTitle)
		if result == nil {
			fmt.Println("Result is nil")
		} else {
			fmt.Printf("%s\n", result.Title)
		}
	} else if *done {
		url := createIssueUrl(*num)
		fmt.Printf("%s\n", *ticketTitle)
		input := createJson(createJsonFormat(*ticketTitle))
		github.UpdateIssues(url, input)
	}
}
*/

func createJsonFormat(str string) string {
	return "{" + str + "}"
}

func createJson(str string) []byte {
	//	input, err := json.Marshal(str)
	input, err := json.MarshalIndent(str, "", "	")
	fmt.Println(input)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return input
}

func createIssueUrl(num int) string {
	return github.IssueUrl + strconv.Itoa(num)
}
