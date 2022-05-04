package register

import (
	"context"
	"fmt"

	"github.com/kiraqjx/x-ray/pd"
	"github.com/kiraqjx/x-ray/proxy"
)

type Register struct {
	pd.UnimplementedXRayServer
}

func (r *Register) Register(ctx context.Context, req *pd.Nodes) (*pd.RegisterResponse, error) {
	if req.Nodes == nil {
		return &pd.RegisterResponse{
			Code: 400,
			Msg:  "nodes can not be null",
		}, nil
	}

	ports := make(map[uint32]string, len(req.Nodes))
	for _, node := range req.Nodes {
		proxyInstance, err := proxy.NewProxy(nil, node)
		if err != nil {
			fmt.Printf("proxy error - ip: %s ; port: %d", node.Ip, node.Port)
			continue
		}
		portChan := make(chan string)
		ctx = context.WithValue(ctx, proxy.PortChanKey, portChan)
		go proxyInstance.Run(ctx)

		port := <-portChan
		ports[node.Port] = port
	}

	return &pd.RegisterResponse{
		Code: 0,
		Msg:  "success",
		Data: ports,
	}, nil
}
