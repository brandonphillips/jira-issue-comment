package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	MethodType          string
	Url                 string
	AuthorizationHeader string
	ContentType         string
	Payload             *strings.Reader
}

func sendComment(environment Config) {
	var request Request
	request = setupRequest(environment)
	if environment.Verbose {
		verboseLogging(environment, request)
	}

	client := &http.Client{}
	req, err := http.NewRequest(request.MethodType, request.Url, request.Payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", request.AuthorizationHeader)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	responseBody := string(body)
	fmt.Println("Response Body\n" + string(responseBody))
	responseJson := map[string]interface{}{}
	json.Unmarshal([]byte(responseBody), &responseJson)
	fmt.Println("\nComment Id: ", responseJson["id"])
}

func setupRequest(environment Config) Request {
	var request Request
	if len(environment.JiraCommentId) > 0 {
		request.MethodType = "PUT"
		request.Url = environment.JiraBaseUrl + "rest/api/2/issue/" + environment.JiraIssueId + "/comment" +
			"/" + environment.JiraCommentId
	} else {
		request.MethodType = "POST"
		request.Url = environment.JiraBaseUrl + "rest/api/2/issue/" + environment.JiraIssueId + "/comment"
	}

	request.AuthorizationHeader = "Basic " + b64.StdEncoding.EncodeToString([]byte(environment.JiraUsername+":"+environment.JiraApiKey))
	request.Payload = strings.NewReader(buildCommentBody(environment))

	return request
}

func verboseLogging(environment Config, request Request) {
	fmt.Println("\nVerbose Logging")
	fmt.Println("Base Url: ", environment.JiraBaseUrl)
	fmt.Println("Full Url: " + request.Url)
	fmt.Println("Username: ", environment.JiraUsername)
	fmt.Println("API Key: ", environment.JiraApiKey)
	fmt.Println("Authorization Header: ", request.AuthorizationHeader)
	fmt.Println("Codefresh Build Link: ", environment.BuildLink)
	fmt.Println("Build Status: ", environment.BuildStatus)
	fmt.Println("Build Message: ", environment.BuildMessage)
	fmt.Println("Comment Payload: ", request.Payload)
	fmt.Println()
}

func buildCommentBody(environment Config) string {
	var buffer bytes.Buffer
	buffer.WriteString("{\"body\": \"")
	buffer.WriteString(environment.BuildMessage + "\\n")
	buffer.WriteString("Pipeline: " + environment.PipelineName + "\\n")
	buffer.WriteString("Build Link: " + environment.BuildLink + "\\n")
	buffer.WriteString("Build Status: " + environment.BuildStatus)
	buffer.WriteString("\"}")

	return buffer.String()
}

// Helper method that checks if the value is populated and loops through them to build a string
