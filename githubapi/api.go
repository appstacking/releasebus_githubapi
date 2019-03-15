package githubapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	rootUrl           = "https://api.github.com/repos/"
	latestVersionPath = "/releases/latest"
)

type repo struct {
	LatestVersion string `json:"tag_name"`
	ReleaseDate   string `json:"published_at"`
}

func GetLatestVersion(ownerName, repoName string) (version, date string, err error) {
	url := rootUrl + ownerName + "/" + repoName + latestVersionPath
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := http.StatusText(resp.StatusCode)
		return "", "", errors.New(errMsg)
	}

	var repoData repo
	err = json.NewDecoder(resp.Body).Decode(&repoData)

	return repoData.LatestVersion, repoData.ReleaseDate, err
}
