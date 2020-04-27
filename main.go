package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
)

type Config struct {
	JiraBaseUrl		string
	Username		string
	ApiKey			string
	JiraIssueId		string
	BuildLink		string
	BuildStatus		string
	BuildMessage	string
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

	return environment
}

func main() {
	// Grab the base path from the environment
	var stepEnvironment Config	
	stepEnvironment = setupEnvironment()

	fmt.Println("baseUrl=", stepEnvironment.JiraBaseUrl)

	url := stepEnvironment.JiraBaseUrl + "rest/api/2/issue/42966/comment"
	method := "POST"

	payload := strings.NewReader("{\"body\": \"Test comment\\nTest comment 3\"}")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
	fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic YnJhbmRvbkBjb2RlZnJlc2guaW86TE9NU2JjWlpGNGpPNHhqaTB6VFdERjI0")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "atlassian.xsrf.token=BQ6L-X8LN-389U-ZPP2_5ee9bac05d2843742a3665bb4113e2191f1eae43_lin")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}