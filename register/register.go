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

	for _, node := range req.Nodes {
		proxy, err := proxy.NewProxy(nil, node)
		if err != nil {
			fmt.Printf("proxy error - ip: %s ; port: %d", node.Ip, node.Port)
			continue
		}
		go proxy.Run(ctx)
	}

	return &pd.RegisterResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}
