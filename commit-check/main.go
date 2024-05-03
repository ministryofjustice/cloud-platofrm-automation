package main

import (
	"fmt"
	"ministryofjustice/cloud-platform-automation/utils"
	"os"
	"strconv"
)

var (
	owner = os.Getenv("GITHUB_REPOSITORY_OWNER")
	repo  = os.Getenv("GITHUB_REPOSITORY")
	prnum = os.Getenv("GITHUB_PR_NUMBER")
	file  = os.Getenv("FILE_NAME")
)

// main function to check if the pull request contains the file with the name specified and no other commited files.
func main() {

	// convert string to int
	pr, err := strconv.Atoi(prnum)
	if err != nil {
		panic(err)
	}

	f, _, err := utils.GetPullRequestFiles(owner, repo, pr)
	if err != nil {
		panic(err)
	}

	b := utils.Files(file, f)
	if b {
		fmt.Printf("pull request commit contains the file with the name %v and no other commited files. This is valid for auto-approval.", file)
		fmt.Println("::set-output name=approval::true")
	}
	if !b {
		fmt.Printf("pull request commit does not contain the file with the name %v or contains other commited files. This is not valid for auto-approval.", file)
		fmt.Println("::set-output name=approval::false")
	}
}
