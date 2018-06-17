package repo

import (
	"fmt"
	"strings"
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
}

func (release *release) String() string {
	var output string

	simpleDateString := release.PublishedAt.Format("2.Jan.2006")
	hasTagOrDate := (strings.Contains(release.Name, release.TagName)) || (strings.Contains(release.Name, simpleDateString))

	if hasTagOrDate {
		output = fmt.Sprintf("---\n### %s\n##### Released on (_%s_) by _%s_\n%s",
			release.TagName,
			release.PublishedAt.Format("Mon 2.Jan.2006"),
			release.Author.Login,
			release.Body)
	} else {
		output = fmt.Sprintf("---\n### %s\n###### Released on (_%s_) by _%s_\n#### %s\n%s",
			release.TagName,
			release.PublishedAt.Format("Mon 2.Jan.2006"),
			release.Author.Login,
			release.Name,
			release.Body)
	}

	return output
}
