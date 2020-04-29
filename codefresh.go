package main

import (
	"fmt"
	"log"
	"os"
)

func createBuildAnnotation(environment Config) {

}

func exportCommentIdVariable(environment Config) {
	file, err := os.OpenFile("/codefresh/volume/env_vars_to_export", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		file.WriteString("JIRA_COMMENT_ID=" + environment.JiraCommentId)
	}
	defer file.Close()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	paths, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(paths)
}
