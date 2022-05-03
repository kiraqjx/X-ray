package proxy

import (
	"context"
	"errors"

	"github.com/kiraqjx/x-ray/pd"
)

type Proxy interface {
	Run(ctx context.Context)
}

func NewProxy(source *pd.Node, target *pd.Node) (Proxy, error) {
	if target.Protocol == pd.Protocol_TCP {
		return &TcpProxy{
			source: source,
			target: target,
		}, nil
	}
	return nil, errors.New("temporary does not support")
}
