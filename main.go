package main

import (
	"fmt"
	"github.com/mattes/go-asciibot"
	"log"
	"net"
	"net/rpc"

	"github.com/appstacking/releasebus_githubapi/rpcservice"
)

func main() {
	// 注册 RPC 服务
	err := rpcservice.RegisterGithubApiService(new(rpcservice.GithubApiService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":12001")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println(asciibot.Random())
	log.Println("server started.enjoy!")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go rpc.ServeConn(conn)
	}
}
