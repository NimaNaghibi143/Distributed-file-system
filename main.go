package main

import (
	"log"

	"github.com/NimaNaghibi143/Distributed-file-system/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tr := p2p.NewTcpTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

	// fmt.Println("Hello babay this is going to be a fantastic project")
}
