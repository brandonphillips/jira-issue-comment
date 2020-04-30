package main

import (
	"fmt"
	"os"
)

func createBuildAnnotation(environment Config) {

}

func exportCommentIdVariable(environment Config) {
	fmt.Println("environment.JiraCommentId: " + environment.JiraCommentId)

	if fileExists(environment.CodefreshVolumePath + "/env_vars_to_export") {
		f, err := os.OpenFile(environment.CodefreshVolumePath+"/env_vars_to_export", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Error opening env_vars_to_export file")
			panic(err)
		} else {
			defer f.Close()

			if _, err = f.WriteString("JIRA_COMMENT_ID=" + environment.JiraCommentId + "\n"); err != nil {
				fmt.Println("Error writing JIRA_COMMENT_ID to env_vars_to_export file")
				panic(err)
			}
		}
	} else {
		fmt.Println("File:" + environment.CodefreshVolumePath + "/env_vars_to_export" +
			"doesn't exist. Unable to write build variable JIRA_COMMENT_ID")
	}

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
