package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

// we could return the Transport interface but then we are testing we need to casttype
// the instance of it.
func NewTcpTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}
