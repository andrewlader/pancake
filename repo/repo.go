package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// GetChangelog returns a markdown representation of the release notes for the provided repo
func GetChangelog(token string, organization string, repo string) (string, int, error) {
	releases, err := getReleasesFromGithub(token, organization, repo)
	if err != nil {
		return "", 0, err
	}

	return processReleaseNotes(releases, organization, repo)
}

func getReleasesFromGithub(token string, organization string, repo string) ([]*release, error) {
	const apiFormat = "https://api.github.com/repos/%s/%s/releases"
	const headerAuth = "Authorization"
	const headerAuthFormat = "token %s"

	api := fmt.Sprintf(apiFormat, organization, repo)
	header := fmt.Sprintf(headerAuthFormat, token)

	request, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(headerAuth, header)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	releaseNotes := make([]*release, 0)
	json.NewDecoder(response.Body).Decode(&releaseNotes)

	if len(releaseNotes) < 1 {
		return nil, errors.New("no release notes found")
	}

	return releaseNotes, nil
}

func processReleaseNotes(releases []*release, organization string, repo string) (string, int, error) {
	waiter := &sync.WaitGroup{}
	for _, release := range releases {
		// have each release generate their release notes in markdown concurrently
		go release.generateMarkdown(waiter, organization, repo)
	}

	numberOfReleaseNotes := len(releases)
	releaseNotes := initializeReleaseNotes(repo, numberOfReleaseNotes)

	waiter.Wait()
	for _, release := range releases {
		releaseNotes = append(releaseNotes, release.markdown)
	}

	changeLog := strings.Join(releaseNotes, "\n")

	return changeLog, numberOfReleaseNotes, nil
}

func initializeReleaseNotes(repo string, numberOfReleaseNotes int) []string {
	heading := fmt.Sprintf("## Changelog for %s", repo)
	autoGeneratedBy := fmt.Sprintf("###### _Autogenerated by_ [pancake](https://github.com/andrewlader/pancake) on %s",
		time.Now().Format("Mon 2.Jan.2006"))
	numberOfReleases := fmt.Sprintf("###### _Found %d release notes_", numberOfReleaseNotes)

	releaseNotes := make([]string, 0)

	releaseNotes = append(releaseNotes, heading)
	releaseNotes = append(releaseNotes, autoGeneratedBy)
	releaseNotes = append(releaseNotes, numberOfReleases)

	return releaseNotes
}
