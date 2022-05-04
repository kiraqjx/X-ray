package proxy

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/kiraqjx/x-ray/pd"
)

type TcpProxy struct {
	source *pd.Node
	target *pd.Node
}

func (t *TcpProxy) Run(ctx context.Context) {
	t.runSourceTcpListener(ctx)
}

// run source listener
func (t *TcpProxy) runSourceTcpListener(ctx context.Context) {
	portChan := ctx.Value(PortChanKey).(chan string)
	if t.source == nil {
		addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:0")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.source = &pd.Node{}
		t.source.Ip = "127.0.0.1"
		t.source.Port = uint32(addr.Port)
	}
	sourceHost := fmt.Sprintf("%s:%d", t.source.Ip, t.source.Port)
	targetHost := fmt.Sprintf("%s:%d", t.target.Ip, t.target.Port)
	sourceListener, err := net.Listen("tcp", sourceHost)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sourceListener.Close()

	portChan <- sourceListener.Addr().String()

	fmt.Printf("Proxy from %s to %s \n", sourceListener.Addr().String(), targetHost)

	for {
		conn, err := sourceListener.Accept()
		if err != nil {
			fmt.Printf("accept error: %v \n", err)
			return
		}
		go t.runTargetTcpSender(conn)
	}
}

// copy source to target
func (t *TcpProxy) runTargetTcpSender(sourceConn net.Conn) {
	defer sourceConn.Close()
	targetSender, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.target.Ip, t.target.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer targetSender.Close()
	io.Copy(sourceConn, targetSender)
}
