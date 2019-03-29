package githubapi

import (
	"net/http"
	"testing"
)

// 测试获取 repo 的最新版本
func TestGetLatestVersion(t *testing.T) {
	t.Run("testGetLatestVersionNormal", testGetLatestVersionNormal)
	t.Run("testGetLatestVersionRepoNotFound", testGetLatestVersionRepoNotFound)
}

// 测试正常获取 repo 的最新版本的情况
func testGetLatestVersionNormal(t *testing.T) {
	testData := []struct {
		owner       string
		repo        string
		wantVersion string
		wantDate    string
	}{
		{
			owner:       "xiaosongfu",
			repo:        "Url2IOElastic",
			wantVersion: "v2.0.0",
			wantDate:    "2016-12-17T16:02:52Z",
		},
		{
			owner:       "xiaosongfu",
			repo:        "asciibot",
			wantVersion: "v1.0.1",
			wantDate:    "2019-03-26T09:31:47Z",
		},
	}

	for _, d := range testData {
		version, date, err := GetLatestVersion(d.owner, d.repo)
		if err != nil {
			t.Errorf(err.Error())
		}

		if d.wantVersion != version {
			t.Errorf("version: want and actual is not match! want is: %s, actual is: %s", d.wantVersion, version)
		}

		if d.wantDate != date {
			t.Errorf("date: want and actual is not match! want is: %s, actual is: %s", d.wantDate, date)
		}
	}
}

// 测试 repo 不存在的情况
func testGetLatestVersionRepoNotFound(t *testing.T) {
	testData := struct {
		owner string
		repo  string
	}{
		owner: "xiaosongfu",
		repo:  "repo_not_found",
	}

	_, _, err := GetLatestVersion(testData.owner, testData.repo)
	if err.Error() != http.StatusText(http.StatusNotFound) {
		t.Errorf("test repo not found case failed")
	}
}
