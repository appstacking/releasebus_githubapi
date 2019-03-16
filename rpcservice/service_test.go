package rpcservice

//import (
//	"net"
//	"net/http"
//	"net/rpc"
//	"testing"
//)
//
//const (
//	rpcNetwork, address = "tcp", ":12008"
//	dialAddr = "localhost" + address
//)
//
//const githubApiServiceFullName  = "com.xiaosongfu.rb.svc.githubapi.LatestRepoVersion"
//
//var listener net.Listener
//
//// 启动 rpc 服务
//func startRpcServer() {
//	RegisterGithubApiService(new(GithubApiService))
//
//	listener, err := net.Listen(rpcNetwork, address)
//	if err != nil {
//		panic(err)
//	}
//
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			panic(err)
//		}
//
//		go rpc.ServeConn(conn)
//	}
//}
//
//// 关闭 rpc 服务
//func stopRpcServer() {
//	listener.Close()
//}
//
//func TestMain(m *testing.M) {
//	// 启动 rpc 服务
//	startRpcServer()
//
//	m.Run()
//
//	// 关闭 rpc 服务
//	stopRpcServer()
//
//}
//
//func TestLatestRepoVersionNormal(t *testing.T) {
//	in := RepoParam{
//		Owner: "xiaosongfu",
//		Repo: "Url2IOElastic",
//		LatestVersion: "v2.0.0",
//		ReleaseDate: "2016-12-17T16:02:52Z",
//	}
//
//	var out RepoParam
//
//	client, err := rpc.Dial(rpcNetwork, dialAddr)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	defer client.Close()
//
//	err = client.Call(githubApiServiceFullName, &in, &out)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if in.Owner != out.Owner || in.Repo == out.Owner || in.LatestVersion != out.LatestVersion || in.ReleaseDate != out.ReleaseDate {
//		t.Fatal()
//	}
//}
//
//func TestLatestRepoVersionRepoNotFound(t *testing.T) {
//	in := RepoParam{
//		Owner: "xiaosongfu",
//		Repo: "not-found-repo",
//	}
//
//	var out RepoParam
//
//	client, err := rpc.Dial(rpcNetwork, dialAddr)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	defer client.Close()
//
//	err = client.Call(serviceName, &in, &out)
//	if err == nil {
//		t.Fatal("test repo not found fail")
//	} else {
//		if err.Error() != http.StatusText(http.StatusOK) {
//			t.Errorf("error is not 'Not Found',test fail, error is:%s", err)
//		}
//	}
//}
