package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"os"
	"os/signal"
	"syscall"
)

var (
	conf = config{}
)

type config struct {
	Command *string
	Client  Client
}

type Client struct {
	Addr *string
}

func main() {
	conf.Command = flag.String("command", "server", "Command:server|client")
	conf.Client.Addr = flag.String("client.addr", "", "peer addr")
	flag.Parse()
	if !(*conf.Command == "server" || *conf.Command == "client") {
		fmt.Println("The command value: server|client")
		return
	}
	var (
		node host.Host
		err  error
	)
	switch *conf.Command {
	case "server":
		node, err = pingServer()
	case "client":
		node, err = pingClient()
	}
	if err != nil {
		panic(err)
	}
	if node == nil {
		return
	}
	if *conf.Command == "server" {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		fmt.Println("Shutting Down...")
	}
	if err := node.Close(); err != nil {
		panic(err)
	}
}

func pingServer() (node host.Host, err error) {
	node, err = libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
		libp2p.Ping(false),
	)
	if err != nil {
		return
	}

	// Configure our own ping protocol
	pingService := &ping.PingService{Host: node}
	node.SetStreamHandler(ping.ID, pingService.PingHandler)

	// print the node's PeerInfo in multiaddr format
	peerInfo := peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	fmt.Println("libp2p node address:", addrs[0])

	return
}
func pingClient() (node host.Host, err error) {
	addr := *conf.Client.Addr

	peer, err := peer.AddrInfoFromString(addr)
	if err != nil {
		return
	}

	node, err = libp2p.New()
	if err != nil {
		return
	}
	if err := node.Connect(context.Background(), *peer); err != nil {
		return node, err
	}

	fmt.Println("Send 5 ping messages to", addr)
	pingService := &ping.PingService{Host: node}
	ch := pingService.Ping(context.Background(), peer.ID)
	for i := 0; i < 5; i++ {
		res := <-ch
		fmt.Println("Pinged", addr, "in", res.RTT)
	}
	return
}
