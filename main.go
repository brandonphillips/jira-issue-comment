package main

import "fmt"

func main() {
	// Grab the environment variables from the step
	var environment Config
	environment = setupEnvironment()

	// Call to add or update a comment
	fmt.Println("main.go - before - environment.JiraCommentId: " + environment.JiraCommentId)
	environment.JiraCommentId = sendComment(environment)
	fmt.Println("main.go - after - environment.JiraCommentId: " + environment.JiraCommentId)
	createBuildAnnotation(environment)
	exportCommentIdVariable(environment)
}
