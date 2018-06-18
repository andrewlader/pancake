package repo

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func Test_HappyCase_Success(t *testing.T) {
	const expectedNumberOfReleases = 14

	jsonTestData := readTestData(t)

	releaseNotes := make([]*release, 0)
	err := json.Unmarshal([]byte(jsonTestData), &releaseNotes)
	if err != nil {
		t.Fatalf("failed to process JSON test data: %v", err)
	}

	releases, actualNumberOfRelease, err := processReleaseNotes(releaseNotes, "some-org", "service-repo")
	if err != nil {
		t.Fatalf("failed to process the releases into markdown: %v", err)
	}

	if actualNumberOfRelease != expectedNumberOfReleases {
		t.Fatalf("failed to get the right number of releases markdown: %v", err)
	}

	if len(releases) == 0 {
		t.Fatalf("failed to generate the releases markdown: %v", err)
	}
}

func readTestData(t *testing.T) []byte {
	const testDataFilename = "./test_data.json"

	jsonTestData, err := ioutil.ReadFile(testDataFilename)
	if err != nil {
		t.Fatalf("failed to read in the test data: %v", err)
	}

	return jsonTestData
}
