package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createBuildAnnotation(environment Config) {

}

func exportCommentIdVariable(environment Config) {
	// dat, err := ioutil.ReadFile("/codefresh/volume/env_vars_to_export")
	// check(err)
	// fmt.Print(string(dat))
	// fmt.Println()

	// file, err := os.OpenFile("/codefresh/volume/env_vars_to_export", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	file.WriteString("JIRA_COMMENT_ID=" + environment.JiraCommentId)
	// }
	// defer file.Close()
	fmt.Println("JiraCommentId: " + environment.JiraCommentId)

	// d1 := []byte("JIRA_COMMENT_ID=" + environment.JiraCommentId)
	// d1 := []byte("JIRA_COMMENT_ID=30812")
	// err2 := ioutil.WriteFile("/codefresh/volume/env_vars_to_export", d1, 0644)
	// check(err2)
	f, err := os.OpenFile("/codefresh/volume/env_vars_to_export", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("JIRA_COMMENT_ID=30812\n"); err != nil {
		panic(err)
	}

	dat2, err := ioutil.ReadFile("/codefresh/volume/env_vars_to_export")
	check(err)
	fmt.Print(string(dat2))
	fmt.Println()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
