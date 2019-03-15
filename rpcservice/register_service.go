package rpcservice

import (
	"net/rpc"
)

const serviceName = "com.xiaosongfu.rb.svc.githubapi"

type GithubApiServiceInterface interface {
	LatestRepoVersion(*RepoParam, *RepoParam) error
}

func RegisterGithubApiService(s GithubApiServiceInterface) error {
	return rpc.RegisterName(serviceName, s)
}
