package rpcservice

import "github.com/appstacking/releasebus_service_githubapi/githubapi"

type GithubApiService struct{}

func (g *GithubApiService) LatestRepoVersion(in *RepoParam, out *RepoParam) error {
	version, date, err := githubapi.GetLatestVersion(in.Owner, in.Repo)
	if err != nil {
		return err
	}

	out.LatestVersion = version
	out.ReleaseDate = date
	return nil
}
