package main

import (
	"fmt"
	"log"

	"github.com/NimaNaghibi143/Distributed-file-system/p2p"
)

func OnPeer(p2p.Peer) error {
	return fmt.Errorf("failed the onpeer func")
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTcpTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

	// fmt.Println("Hello babay this is going to be a fantastic project")
}
