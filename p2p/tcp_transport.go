package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established Connection.
type TCPPeer struct {
	// conn is the underlying connection of the peer.
	conn net.Conn
	// if we dial and retrieve a connection => outbound == true
	// if we accept and retrieve a connection => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Close implements the peer interface
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rcpch    chan RPC

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

// we could return the Transport interface but then we are testing we need to casttype
// the instance of it.
func NewTcpTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rcpch:            make(chan RPC),
	}
}

// Consume implements the transport interface, which will return read-only channel
// for reading the incoming messages recieved from another peer in the network.
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rcpch
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		fmt.Printf("New incoming connection %+v\n", conn)
		// keep looping and keep listening
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	// Read loop
	rpc := RPC{}

	// buf := make([]byte, 2000)

	for {

		// n, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Printf("TCP error: %s\n", err)
		// }

		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

		rpc.From = conn.RemoteAddr()
		t.rcpch <- rpc

		fmt.Printf("message: %+v\n", rpc)
	}
}
