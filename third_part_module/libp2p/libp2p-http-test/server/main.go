package main

import (
	"fmt"
	"github.com/libp2p/go-libp2p"
	gostream "github.com/libp2p/go-libp2p-gostream"
	p2phttp "github.com/libp2p/go-libp2p-http"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	myProtocol protocol.ID = "/testiti-test"
)

func main() {
	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
	)
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 1 {
		httpServer(host)
	} else {
		httpClient(host, os.Args[1])
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("System shutting down ...")
	if err := host.Close(); err != nil {
		panic(err)
	}
}

func httpServer(host host.Host) {
	listener, _ := gostream.Listen(host, myProtocol)
	defer listener.Close()
	go func() {
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hi!"))
		})
		server := &http.Server{}
		server.Serve(listener)
	}()
	peerInfo := peer.AddrInfo{
		host.ID(), host.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("httpServer address:", addrs[0])
}

func httpClient(host host.Host, addr string) {
	addrInfo, err := peer.AddrInfoFromString(addr)
	if err != nil {
		panic(err)
	}
	host.Peerstore().AddAddrs(addrInfo.ID, addrInfo.Addrs, peerstore.PermanentAddrTTL)

	tr := &http.Transport{}
	tr.RegisterProtocol("libp2p", p2phttp.NewTransport(host, p2phttp.ProtocolOption(myProtocol)))
	client := &http.Client{Transport: tr}

	res, err := client.Get(fmt.Sprintf("libp2p://%s/hello", addrInfo.ID.String()))
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %#v", res)
}
