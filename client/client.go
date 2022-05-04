package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/kiraqjx/x-ray/config"
	"github.com/kiraqjx/x-ray/pd"
	"gopkg.in/yaml.v3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientConfigPath string

var clientConfig config.ClientConfig

func init() {
	flag.StringVar(&clientConfigPath, "client-config", "../config/client-config.yaml", "the client config file path")
	flag.Parse()
	file, err := os.Open(clientConfigPath)
	if err != nil {
		panic(err)
	}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&clientConfig)
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := grpc.Dial(clientConfig.GrpcServer.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := pd.NewXRayClient(conn)

	node := &pd.Node{
		Protocol: pd.Protocol(clientConfig.ProxyHosts[0].Protocol),
		Ip:       clientConfig.ProxyHosts[0].Host,
		Port:     clientConfig.ProxyHosts[0].Port,
	}

	nodes := make([]*pd.Node, 1)
	nodes[0] = node

	nodesPackage := &pd.Nodes{Nodes: nodes}

	rsp, err := client.Register(context.TODO(), nodesPackage)

	if err != nil {
		fmt.Println(err)
		return
	}

	printProxyInfo(rsp.Data)

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", clientConfig.ProxyHosts[0].Host, clientConfig.ProxyHosts[0].Port))
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

func printProxyInfo(data map[uint32]string) {
	for key, value := range data {
		fmt.Printf("%d -> %s", key, value)
	}
	fmt.Print("\n")
}
