package main

import (
	"fmt"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {
	// Grab the base path from the environment
	baseUrl := "https://jira.atlassian.com"
	passedUrl, exists := os.LookupEnv("baseUrl")

	if exists {
		fmt.Println(baseUrl)
		baseUrl = passedUrl
	}

	url := baseUrl + "rest/api/2/issue/42966/comment"
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