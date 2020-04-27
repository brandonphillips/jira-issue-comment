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
	Verbose       bool
}

func setupEnvironment() Config {
	var environment Config

	// Setting default log level
	verbose, err := strconv.ParseBool(getEnvironmentVariable("verbose", false))
	if err != nil {
		environment.Verbose = false
	} else {
		environment.Verbose = verbose
	}

	// Verify that the required variables are passed in with the step
	environment.JiraBaseUrl = getEnvironmentVariable("baseUrl", true)
	environment.JiraUsername = getEnvironmentVariable("username", true)
	environment.JiraApiKey = getEnvironmentVariable("apiKey", true)
	environment.JiraIssueId = getEnvironmentVariable("issue", true)

	// Codefresh provided variables
	environment.BuildLink = getEnvironmentVariable("buildLink", false)
	environment.BuildStatus = getEnvironmentVariable("buildStatus", false)
	environment.BuildMessage = getEnvironmentVariable("buildMessage", false)

	// Comment id for updating build status of original comment - won't be provided for initial run
	environment.JiraCommentId = getEnvironmentVariable("commentId", false)

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
