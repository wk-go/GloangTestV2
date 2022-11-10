package main

import (
	"fmt"
	"github.com/libp2p/go-libp2p"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	node, err := libp2p.New()

	if err != nil {
		panic(err)
	}

	fmt.Println("Listen Address:", node.Addrs())
	if err := node.Close(); err != nil {
		panic(err)
	}

	node2, err2 := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"))
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Nod2 Listen Address:", node2.Addrs())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received Signal Shutting down ...")
	if err := node2.Close(); err != nil {
		panic(err)
	}
}
