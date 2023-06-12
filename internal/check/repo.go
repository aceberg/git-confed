package check

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// IsRepo checks if the path is git repo
func IsRepo(path string) bool {

	_, err := os.Stat(path + "/.git")

	return err == nil
}

// ParseConfig - get user and remotes
func ParseConfig(path string) (string, []string) {
	var text, user string
	var remote []string

	file, err := os.Open(path + "/.git/config")
	IfError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = scanner.Text()

		re, _ := regexp.Compile(`name =`)
		if re.FindString(text) != "" {
			user = re.ReplaceAllString(text, "")
		}
		re, _ = regexp.Compile(`\[remote \"`)
		if re.FindString(text) != "" {
			text = re.ReplaceAllString(text, "")
			re, _ = regexp.Compile(`\"\]`)
			text = re.ReplaceAllString(text, "")
			remote = append(remote, text)
		}
	}

	// log.Println("BRANCH =", branch, "USER =", user, "REMOTE =", remote)

	return user, remote
}

// Branch - returns current git branch
func Branch(path string) string {
	var branch string

	file, err := os.ReadFile(path + "/.git/HEAD")
	IfError(err)

	re, _ := regexp.Compile(`ref: refs\/heads\/`)
	branch = re.ReplaceAllString(string(file), "")

	return branch
}

// URL - checks if string is in .git/config
func URL(path string, urls []string) []string {
	var found []string

	file, err := os.ReadFile(path + "/.git/config")
	IfError(err)

	for _, url := range urls {
		if strings.Contains(string(file), url) {
			found = append(found, url)
		}
	}

	return found
}

// ReplaceReponame - replace variable $REPONAME with the name of the repo
func ReplaceReponame(block, name string) string {

	re, _ := regexp.Compile(`\$REPONAME`)
	newBlock := re.ReplaceAllString(block, name)

	return newBlock
}
