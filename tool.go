package ghaction

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GetMatchedFiles(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

func ParseRepoURL(repoURL string) (owner, repo string, err error) {
	strArray := strings.Split(repoURL, "/")
	if len(strArray) < 2 {
		return "", "", fmt.Errorf("could parse repo URL %s", repoURL)
	}
	strLen := len(strArray)
	return strArray[strLen-2], strArray[strLen-1], nil
}
