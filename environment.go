package main

import (
	"log"
	"os"
)

type Config struct {
	JiraBaseUrl		string
	JiraUsername	string
	JiraApiKey		string
	JiraIssueId		string
	BuildLink		string
	BuildStatus		string
	BuildMessage	string
	Verbose			bool
}

func setupEnvironment() Config {
	var environment Config

	// Verify that the Jira base url was passed in with the step
	passedUrl, exists := os.LookupEnv("baseUrl")
	if !exists {
		log.Fatal("Fatal Error - Exiting Step: Jira base url is a required field")		
	} else {
		environment.JiraBaseUrl = passedUrl
	}

	environment.JiraUsername = "brandon@codefresh.io"
	environment.JiraApiKey = "LOMSbcZZF4jO4xji0zTWDF24"
	environment.JiraIssueId = "42966"
	environment.BuildLink = "https://g.codefresh.io/build/5ea4f9bc9365eb3265d5fe97"
	environment.BuildStatus = "Successful"
	environment.BuildMessage = "Some custom message"
	environment.Verbose = true

	return environment
}