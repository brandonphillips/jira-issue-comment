package main

import (
	b64 "encoding/base64"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {
	// Grab the environment variables from the step
	var environment Config	
	environment = setupEnvironment()

	fmt.Println("baseUrl=", environment.JiraBaseUrl)

	url := environment.JiraBaseUrl + "rest/api/2/issue/" + environment.JiraIssueId + "/comment"
	method := "POST"

	payload := strings.NewReader("{\"body\": \"Test comment\\nTest comment 3\"}")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
	fmt.Println(err)
	}

	// Convert the username and the api key to a base64 encoded authorization key
	authorizationHeader := "Basic " + b64.StdEncoding.EncodeToString([]byte(environment.JiraUsername + ":" + environment.JiraApiKey))
	if environment.Verbose {
		fmt.Println("Authorization Header:", authorizationHeader)
	}	
	
	req.Header.Add("Authorization", authorizationHeader)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}