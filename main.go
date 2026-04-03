package main

import (
	"log"

	"github.com/NimaNaghibi143/Distributed-file-system/p2p"
)

func main() {
	tr := p2p.NewTcpTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

	// fmt.Println("Hello babay this is going to be a fantastic project")
}
