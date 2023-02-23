package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

const (
	epicLinkFieldID = ""
	username        = ""
	password        = ""
)

var (
	jql = "issueKey IN (APP-54995,APP-54995,APP-54991,APP-55085,APP-54991,APP-55009,APP-54697,APP-54348,APP-54995,APP-55009,APP-54025,APP-54348,APP-54697,APP-54990,APP-55064,APP-54984,APP-53613,APP-54647,APP-55252,APP-53614,APP-54992,APP-53608,APP-55236,APP-54006,APP-53615,APP-54992,APP-53614,APP-54011,APP-55265,APP-55106,APP-55412,APP-55106,APP-55203,APP-54697,APP-55238,APP-53614,APP-55485,APP-55475,APP-55474,APP-54986,APP-55116,APP-54013,APP-54986,APP-55412,APP-53617,APP-55236,APP-55611,APP-54986,APP-55494,APP-55252,APP-55203,APP-55591,APP-55381,APP-53609,APP-54994,APP-53614,APP-54118,APP-55591,APP-55466,APP-55117)"
)

func getDefaultSearchOption() *jira.SearchOptions {
	return &jira.SearchOptions{
		Fields: []string{"issuetype", "parent", epicLinkFieldID},
	}
}

func main() {
	jiraURL := ""

	jiraAuth := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	client, err := jira.NewClient(jiraAuth.Client(), jiraURL)
	if err != nil {
		panic(err)
	}

	u, _, err := client.Issue.Search(jql, getDefaultSearchOption())
	if err != nil {
		panic(err)
	}

	fmt.Println("Success!", u)
}
