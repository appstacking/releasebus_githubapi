package rpcservice

import (
	"net/rpc"
)

const serviceName = "releasebus.githubapi.service"

type GithubApiServiceInterface interface {
	LatestRepoVersion(*RepoParam, *RepoParam) error
}

func RegisterGithubApiService(s GithubApiServiceInterface) error {
	return rpc.RegisterName(serviceName, s)
}
