package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/kiraqjx/x-ray/config"
	"github.com/kiraqjx/x-ray/pd"
	"github.com/kiraqjx/x-ray/register"
	"gopkg.in/yaml.v3"

	"google.golang.org/grpc"
)

var serverConfigPath string

var serverConfig config.ServerConfig

func init() {
	flag.StringVar(&serverConfigPath, "server-config", "../config/server-config.yaml", "the server config file path")
	flag.Parse()
	file, err := os.Open(serverConfigPath)
	if err != nil {
		panic(err)
	}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&serverConfig)
	if err != nil {
		panic(err)
	}
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", serverConfig.Grpc.Port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pd.RegisterXRayServer(s, &register.Register{})

	defer func() {
		s.Stop()
		listen.Close()
	}()

	fmt.Println("Serving 8081...")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
