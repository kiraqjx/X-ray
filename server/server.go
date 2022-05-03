package main

import (
	"fmt"
	"net"

	"github.com/kiraqjx/x-ray/pd"
	"github.com/kiraqjx/x-ray/register"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8081")
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
