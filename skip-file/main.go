package main

import (
	"fmt"
	"ministryofjustice/cloud-platform-automation/utils"
)

func main() {
	f, _, err := utils.GetPullRequestFiles("ministryofjustice", "cloud-platform-environments", 22597)
	if err != nil {
		panic(err)
	}

	b := utils.Files("APPLY_PIPELINE_SKIP_THIS_NAMESPACE", f)
	if b {
		fmt.Println("pull request commit contains the file with the name 'APPLY_PIPELINE_SKIP_THIS_NAMESPACE' and no other commited files. This is valid for auto-approval.")
		fmt.Println("::set-output name=approval::true")
	}
	if !b {
		fmt.Println("pull request commit does not contain the file with the name 'APPLY_PIPELINE_SKIP_THIS_NAMESPACE' or contains other commited files. This is not valid for auto-approval.")
		fmt.Println("::set-output name=approval::false")
	}
}
