package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	//	now := time.Now()
	loc, _ := time.LoadLocation("UTC")
	onemonthago := time.Date(2016, 2, 1, 0, 0, 0, 0, loc)
	oneyearago := time.Date(2015, 2, 1, 0, 0, 0, 0, loc)

	a := lessthan(result, onemonthago)
	b := lessthan(result, oneyearago)
	c := morethan(result, oneyearago)

	layout := "2006-01-02 15:04:05"

	fmt.Println("Less than one month ============================= ")
	for _, item := range a {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format(layout))
	}
	fmt.Println("Less than one year ============================= ")
	for _, item := range b {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format(layout))
	}

	fmt.Println("More than one year ============================= ")
	for _, item := range c {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format(layout))
	}
	fmt.Println("Done ============================= ")
}

func lessthan(result *github.IssuesSearchResult, term time.Time) map[time.Time]*github.Issue {
	issues := make(map[time.Time]*github.Issue)
	for _, item := range result.Items {
		ctime := item.CreatedAt
		if ctime.After(term) {
			fmt.Println(ctime)
			issues[ctime] = item
		}
	}
	return issues
}

func morethan(result *github.IssuesSearchResult, term time.Time) map[time.Time]*github.Issue {
	issues := make(map[time.Time]*github.Issue)
	for _, item := range result.Items {
		ctime := item.CreatedAt
		if ctime.Before(term) {
			fmt.Println(ctime)
			issues[ctime] = item
		}
	}
	return issues
}
