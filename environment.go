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
	Verbose       bool
	InfoValues    []CommentValue
}

type CommentValue struct {
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

	// Comment id for updating build status of original comment - won't be provided for initial run
	environment.JiraCommentId = getEnvironmentVariable("JIRA_COMMENT_ID", false)

	// Codefresh provided variables
	environment.InfoValues = append(environment.InfoValues,
		CommentValue{"", getEnvironmentVariable("BUILD_MESSAGE", false)},
		CommentValue{"Pipeline: ", getEnvironmentVariable("CF_PIPELINE_NAME", false)},
		CommentValue{"Build Url: ", getEnvironmentVariable("CF_BUILD_URL", false)},
		CommentValue{"Build Status: ", getEnvironmentVariable("BUILD_STATUS", false)},
	)

	if convertStringToBool(getEnvironmentVariable("ADD_BRANCH_INFO", false)) {
		environment = appendCommentValueIfPopulated(environment, "Branch: ", "CF_BRANCH")
	}

	if convertStringToBool(getEnvironmentVariable("ADD_COMMIT_INFO", false)) {
		environment = appendCommentValueIfPopulated(environment, "Commit Author: ", "CF_COMMIT_AUTHOR")
		environment = appendCommentValueIfPopulated(environment, "Commit Message: ", "CF_COMMIT_MESSAGE")
		environment = appendCommentValueIfPopulated(environment, "Commit Url: ", "CF_COMMIT_URL")
	}

	if convertStringToBool(getEnvironmentVariable("ADD_PR_INFO", false)) {
		environment = appendCommentValueIfPopulated(environment, "Pull Request Action: ", "CF_PULL_REQUEST_ACTION")
		environment = appendCommentValueIfPopulated(environment, "Pull Request Target: ", "CF_PULL_REQUEST_TARGET")
		environment = appendCommentValueIfPopulated(environment, "Pull Request Number: ", "CF_PULL_REQUEST_NUMBER")
		environment = appendCommentValueIfPopulated(environment, "Pull Request Id: ", "CF_PULL_REQUEST_ID")
	}

	fmt.Printf("\n%v\n", environment.InfoValues)

	return environment
}

func appendCommentValueIfPopulated(environment Config, displayText string, environmentVariableName string) Config {
	environmentValue := getEnvironmentVariable(environmentVariableName, false)
	if environmentValue != "${{"+environmentVariableName+"}}" {
		environment.InfoValues = append(environment.InfoValues, CommentValue{displayText, environmentValue})
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
