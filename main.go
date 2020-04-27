package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Grab the environment variables from the step
	var environment Config
	environment = setupEnvironment()

	url := environment.JiraBaseUrl + "rest/api/2/issue/" + environment.JiraIssueId + "/comment"
	method := "POST"

	payload := strings.NewReader("{\"body\": \"Test comment\\nTest comment 3\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	// Convert the username and the api key to a base64 encoded authorization key
	authorizationHeader := "Basic " + b64.StdEncoding.EncodeToString([]byte(environment.JiraUsername+":"+environment.JiraApiKey))
	if environment.Verbose {
		fmt.Println("Base Url: ", environment.JiraBaseUrl)
		fmt.Println("Full Url: " + url)
		fmt.Println("Username: ", environment.JiraUsername)
		fmt.Println("API Key: ", environment.JiraApiKey)
		fmt.Println("Authorization Header: ", authorizationHeader)
		fmt.Println("Codefresh Build Link: ", environment.BuildLink)
		fmt.Println("Build Status: ", environment.BuildStatus)
		fmt.Println("Build Message: ", environment.BuildMessage)
	}

	req.Header.Add("Authorization", authorizationHeader)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
