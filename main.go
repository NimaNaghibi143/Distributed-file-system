package main

import (
	"fmt"
	"log"

	"github.com/NimaNaghibi143/Distributed-file-system/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	//fmt.Printf("doing some logic with the peer outside TCPTransport.")
	return nil
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
