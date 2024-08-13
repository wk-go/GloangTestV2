package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	"github.com/multiformats/go-multiaddr"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	service := flag.String("service", "", "type:relay|client|server")
	relayAddress := flag.String("relay-addr", "", "relay address")    // required for client or server
	serverAddress := flag.String("server-addr", "", "server address") // required for server
	flag.Parse()
	var (
		h   host.Host
		err error
		ch  = make(chan os.Signal, 1)
	)
	switch *service {
	case "relay":
		//中继
		h, _, err = NewRelay()
		if err != nil {
			panic(err)
		}
	case "client", "server":
		fallthrough
	default:
		if len(*relayAddress) == 0 {
			panic(fmt.Errorf("relay address is required"))
		}
		relayAddrInfo, err := peer.AddrInfoFromString(*relayAddress)
		if err != nil {
			panic(err)
		}
		h, err = NewNode()
		if err = h.Connect(context.Background(), *relayAddrInfo); err != nil {
			panic(err)
		}
		//服务端
		if *service == "server" {
			h.SetStreamHandler("/protocol01", func(s network.Stream) {
				if _, err := s.Write([]byte("Hello world")); err != nil {
					fmt.Printf("protocol01 err: %s", err)
				}
				s.Close()
			})
			_, err := client.Reserve(context.Background(), h, *relayAddrInfo)
			if err != nil {
				panic(err)
			}
		} else {
			//客户端

			if len(*serverAddress) == 0 {
				panic(fmt.Errorf("server address is required"))
			}
			serverAddrInfo, err := peer.AddrInfoFromString(*serverAddress)
			if err != nil {
				panic(err)
			}
			relayAddr, err := multiaddr.NewMultiaddr("/p2p/" + relayAddrInfo.ID.String() + "/p2p-circuit/p2p/" + serverAddrInfo.ID.String())
			if err != nil {
				panic(err)
			}
			serverRelayAddrInfo := peer.AddrInfo{
				ID:    serverAddrInfo.ID,
				Addrs: []multiaddr.Multiaddr{relayAddr},
			}
			if err := h.Connect(context.Background(), serverRelayAddrInfo); err != nil {
				panic(err)
			}
			s, err := h.NewStream(network.WithUseTransient(context.Background(), "protocol01"), serverAddrInfo.ID, "/protocol01")
			if err != nil {
				panic(err)
			}

			go func() {
				result := make([]byte, 1024*32)
				b := make([]byte, 4096)
				for {
					n, err := s.Read(b)
					if err == io.EOF {
						break
					}
					result = append(result, b[:n]...)
				}
				fmt.Printf("Result: %s\n", result)
			}()
		}

	}

	addrInfo := peer.AddrInfo{
		ID:    h.ID(),
		Addrs: h.Addrs(),
	}

	addrs, err := peer.AddrInfoToP2pAddrs(&addrInfo)
	if err != nil {
		panic(err)
	}

	for i, v := range addrs {
		fmt.Printf("%s address[%d]:%s\n", *service, i, v.String())
	}

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Shutting Down...")
	if err = h.Close(); err != nil {
		fmt.Errorf("err:%s\n", err)
	}
}
