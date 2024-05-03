package utils

import (
	"strings"

	"github.com/google/go-github/v61/github"
)

var (
	skipfile bool
)

func Files(sfname string, files []*github.CommitFile) bool {

	for _, file := range files {
		if strings.Contains(file.GetFilename(), sfname) {
			skipfile = true
		} else {
			return false
		}
	}

	if skipfile {
		if len(files) == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
