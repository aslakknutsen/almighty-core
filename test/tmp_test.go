package test

import (
	"fmt"
	"testing"

	"github.com/andygrunwald/go-jira"
)

func TestMigration(t *testing.T) {
	jiraClient, err := jira.NewClient(nil, "https://issues.jboss.org/")
	if err != nil {
		panic(err)
	}

	res, err := jiraClient.Authentication.AcquireSessionCookie("aslaka", "tuXhuset1k")
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	issue, _, err := jiraClient.Issue.Get("ARQ-2024")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}
