package main

import (
	"fmt"
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
	InfoValues    []OptionalInfoValue
}

type OptionalInfoValue struct {
	DisplayText string
	Value       string
}

func setupEnvironment() Config {
	var environment Config

	environment.Verbose = convertStringToBool(getEnvironmentVariable("VERBOSE", false))

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

	if convertStringToBool(getEnvironmentVariable("ADD_BRANCH_INFO", false)) {
		environment.InfoValues = append(environment.InfoValues,
			OptionalInfoValue{
				DisplayText: "Branch: ",
				Value:       getEnvironmentVariable("CF_BRANCH", false),
			})
	}

	if convertStringToBool(getEnvironmentVariable("ADD_COMMIT_INFO", false)) {
		fmt.Println("Add commit info is set")
	}

	if convertStringToBool(getEnvironmentVariable("ADD_PR_INFO", false)) {
		fmt.Println("Add pr info is set")
	}

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

func convertStringToBool(stringValue string) bool {
	boolValue, err := strconv.ParseBool(stringValue)
	if err != nil {
		boolValue = false
	}
	return boolValue
}
