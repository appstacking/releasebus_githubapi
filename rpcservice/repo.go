package rpcservice

type RepoParam struct {
	Id            int
	Owner         string
	Repo          string
	LatestVersion string
	ReleaseDate   string
}
