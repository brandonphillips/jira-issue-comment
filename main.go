package main

func main() {
	// Grab the environment variables from the step
	var environment Config
	environment = setupEnvironment()

	// Call to add or update a comment
	sendComment(environment)
	createBuildAnnotation(environment)
	exportCommentIdVariable(environment)
}
