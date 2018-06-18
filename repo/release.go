package repo

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
)

type release struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Draft   bool   `json:"draft"`
	Author  struct {
		Login string `json:"login"`
	} `json:"author"`
	Prerelease  bool      `json:"prerelease"`
	PublishedAt time.Time `json:"published_at"`
	Body        string    `json:"body"`
	markdown    string
}

func (release *release) generateMarkdown(waiter *sync.WaitGroup, organization string, repo string) {
	const formatSimpleDate = "2.Jan.2006"
	const formatSimpleDateWithDay = "Mon 2.Jan.2006"
	var output string

	waiter.Add(1)
	defer waiter.Done()

	simpleDateString := release.PublishedAt.Format(formatSimpleDate)
	hasTagOrDate := (strings.Contains(release.Name, release.TagName)) || (strings.Contains(release.Name, simpleDateString))

	release.replaceSpecialGithubPullRequests(organization, repo)
	release.replaceRegularGithubPullRequests(organization, repo)

	if hasTagOrDate {
		output = fmt.Sprintf("---\n### %s\n##### Released on (_%s_) by _%s_\n%s",
			release.TagName,
			release.PublishedAt.Format(formatSimpleDateWithDay),
			release.Author.Login,
			release.Body)
	} else {
		output = fmt.Sprintf("---\n### %s\n###### Released on (_%s_) by _%s_\n#### %s\n%s",
			release.TagName,
			release.PublishedAt.Format(formatSimpleDateWithDay),
			release.Author.Login,
			release.Name,
			release.Body)
	}

	release.markdown = output
}

func (release *release) replaceSpecialGithubPullRequests(organization string, repo string) {
	regex := regexp.MustCompile(`(\[GH\-)([0-9]+)\]`)
	substitution := fmt.Sprintf(`[[GH-$2]](https://github.com/%s/%s/pull/$2)`, organization, repo)
	release.Body = regex.ReplaceAllString(release.Body, substitution)
}

func (release *release) replaceRegularGithubPullRequests(organization string, repo string) {
	regex := regexp.MustCompile(`([^-])#([0-9]+)`)
	substitution := fmt.Sprintf(`$1[[GH-$2]](https://github.com/%s/%s/pull/$2)`, organization, repo)
	release.Body = regex.ReplaceAllString(release.Body, substitution)
}
