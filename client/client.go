package main

import (
	"context"
	"fmt"
	"net"

	"github.com/kiraqjx/x-ray/pd"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serviceHost := "127.0.0.1:8081"
	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := pd.NewXRayClient(conn)

	node := &pd.Node{
		Ip:   "127.0.0.1",
		Port: 8082,
	}

	nodes := make([]*pd.Node, 1)
	nodes[0] = node

	nodesPackage := &pd.Nodes{Nodes: nodes}

	rsp, err := client.Register(context.TODO(), nodesPackage)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)

	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("hello world")
	conn.Close()
}
