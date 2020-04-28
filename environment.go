package main

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	JiraBaseUrl   string
	JiraUsername  string
	JiraApiKey    string
	JiraIssueId   string
	JiraCommentId string
	BuildLink     string
	BuildStatus   string
	BuildMessage  string
	PipelineName  string
	Verbose       bool
}

func setupEnvironment() Config {
	var environment Config

	// Setting default log level
	verbose, err := strconv.ParseBool(getEnvironmentVariable("VERBOSE", false))
	if err != nil {
		environment.Verbose = false
	} else {
		environment.Verbose = verbose
	}

	// Verify that the required variables are passed in with the step
	environment.JiraBaseUrl = getEnvironmentVariable("JIRA_BASE_URL", true)
	environment.JiraUsername = getEnvironmentVariable("JIRA_USERNAME", true)
	environment.JiraApiKey = getEnvironmentVariable("JIRA_API_KEY", true)
	environment.JiraIssueId = getEnvironmentVariable("JIRA_ISSUE", true)

	// Codefresh provided variables
	environment.BuildLink = getEnvironmentVariable("CF_BUILD_URL", false)
	environment.BuildStatus = getEnvironmentVariable("BUILD_STATUS", false)
	environment.BuildMessage = getEnvironmentVariable("BUILD_MESSAGE", false)
	environment.PipelineName = getEnvironmentVariable("CF_PIPELINE_NAME", false)

	// Comment id for updating build status of original comment - won't be provided for initial run
	environment.JiraCommentId = getEnvironmentVariable("JIRA_COMMENT_ID", false)

	return environment
}

func getEnvironmentVariable(environmentVariable string, required bool) string {
	desiredVariable, exists := os.LookupEnv(environmentVariable)
	if !exists && required {
		log.Fatalf("Fatal Error - Exiting Step: %s is a required field", environmentVariable)
	} else if !exists {
		log.Printf("Environment variable %s not set", environmentVariable)
	}

	return desiredVariable
}
